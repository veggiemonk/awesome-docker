package tui

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"unicode/utf8"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/veggiemonk/awesome-docker/internal/cache"
)

type panel int

const (
	panelTree panel = iota
	panelList
)

const entryHeight = 5 // lines rendered per entry in the list panel
const scrollOff = 4   // minimum lines/entries kept visible above and below cursor

// Model is the top-level Bubbletea model.
type Model struct {
	roots    []*TreeNode
	flatTree []FlatNode

	activePanel    panel
	treeCursor     int
	treeOffset     int
	listCursor     int
	listOffset     int
	currentEntries []cache.HealthEntry

	filtering  bool
	filterText string

	width, height int
}

// New creates a new Model from health cache entries.
func New(entries []cache.HealthEntry) Model {
	roots := BuildTree(entries)
	// Expand first root by default
	if len(roots) > 0 {
		roots[0].Expanded = true
	}
	flat := FlattenVisible(roots)

	m := Model{
		roots:    roots,
		flatTree: flat,
	}
	m.updateCurrentEntries()
	return m
}

// Init returns an initial command.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case openURLMsg:
		return m, nil

	case tea.KeyPressMsg:
		// Filter mode input
		if m.filtering {
			return m.handleFilterKey(msg)
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			if m.activePanel == panelTree {
				m.activePanel = panelList
			} else {
				m.activePanel = panelTree
			}
		case "/":
			m.filtering = true
			m.filterText = ""
		default:
			if m.activePanel == panelTree {
				return m.handleTreeKey(msg)
			}
			return m.handleListKey(msg)
		}
	}
	return m, nil
}

func (m Model) handleFilterKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.filtering = false
		m.filterText = ""
		m.flatTree = FlattenVisible(m.roots)
		m.updateCurrentEntries()
	case "enter":
		m.filtering = false
	case "backspace":
		if len(m.filterText) > 0 {
			m.filterText = m.filterText[:len(m.filterText)-1]
			m.applyFilter()
		}
	default:
		r := msg.String()
		if utf8.RuneCountInString(r) == 1 {
			m.filterText += r
			m.applyFilter()
		}
	}
	return m, nil
}

func (m *Model) applyFilter() {
	if m.filterText == "" {
		m.flatTree = FlattenVisible(m.roots)
		m.updateCurrentEntries()
		return
	}

	query := strings.ToLower(m.filterText)
	var filtered []cache.HealthEntry
	for _, root := range m.roots {
		for _, e := range root.AllEntries() {
			if strings.Contains(strings.ToLower(e.Name), query) ||
				strings.Contains(strings.ToLower(e.Description), query) ||
				strings.Contains(strings.ToLower(e.Category), query) {
				filtered = append(filtered, e)
			}
		}
	}
	m.currentEntries = filtered
	m.listCursor = 0
	m.listOffset = 0
}

func (m Model) handleTreeKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.treeCursor > 0 {
			m.treeCursor--
			m.adjustTreeScroll()
			m.updateCurrentEntries()
		}
	case "down", "j":
		if m.treeCursor < len(m.flatTree)-1 {
			m.treeCursor++
			m.adjustTreeScroll()
			m.updateCurrentEntries()
		}
	case "enter", " ":
		if m.treeCursor < len(m.flatTree) {
			node := m.flatTree[m.treeCursor].Node
			if node.HasChildren() {
				node.Expanded = !node.Expanded
				m.flatTree = FlattenVisible(m.roots)
				if m.treeCursor >= len(m.flatTree) {
					m.treeCursor = len(m.flatTree) - 1
				}
			}
			m.adjustTreeScroll()
			m.updateCurrentEntries()
		}
	case "ctrl+d", "pgdown":
		half := m.treePanelHeight() / 2
		if half < 1 {
			half = 1
		}
		m.treeCursor += half
		if m.treeCursor >= len(m.flatTree) {
			m.treeCursor = len(m.flatTree) - 1
		}
		m.adjustTreeScroll()
		m.updateCurrentEntries()
	case "ctrl+u", "pgup":
		half := m.treePanelHeight() / 2
		if half < 1 {
			half = 1
		}
		m.treeCursor -= half
		if m.treeCursor < 0 {
			m.treeCursor = 0
		}
		m.adjustTreeScroll()
		m.updateCurrentEntries()
	case "g", "home":
		m.treeCursor = 0
		m.adjustTreeScroll()
		m.updateCurrentEntries()
	case "G", "end":
		m.treeCursor = len(m.flatTree) - 1
		m.adjustTreeScroll()
		m.updateCurrentEntries()
	case "right", "l":
		if m.treeCursor < len(m.flatTree) {
			node := m.flatTree[m.treeCursor].Node
			if node.HasChildren() && !node.Expanded {
				node.Expanded = true
				m.flatTree = FlattenVisible(m.roots)
				m.adjustTreeScroll()
				m.updateCurrentEntries()
			} else {
				m.activePanel = panelList
			}
		}
	case "left", "h":
		if m.treeCursor < len(m.flatTree) {
			node := m.flatTree[m.treeCursor].Node
			if node.HasChildren() && node.Expanded {
				node.Expanded = false
				m.flatTree = FlattenVisible(m.roots)
				m.adjustTreeScroll()
				m.updateCurrentEntries()
			}
		}
	}
	return m, nil
}

func (m *Model) adjustTreeScroll() {
	visible := m.treePanelHeight()
	off := scrollOff
	if off > visible/2 {
		off = visible / 2
	}
	if m.treeCursor < m.treeOffset+off {
		m.treeOffset = m.treeCursor - off
	}
	if m.treeCursor >= m.treeOffset+visible-off {
		m.treeOffset = m.treeCursor - visible + off + 1
	}
	if m.treeOffset < 0 {
		m.treeOffset = 0
	}
}

func (m Model) treePanelHeight() int {
	h := m.height - 6 // header, footer, borders, title
	if h < 1 {
		h = 1
	}
	return h
}

func (m Model) handleListKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.listCursor > 0 {
			m.listCursor--
			m.adjustListScroll()
		}
	case "down", "j":
		if m.listCursor < len(m.currentEntries)-1 {
			m.listCursor++
			m.adjustListScroll()
		}
	case "ctrl+d", "pgdown":
		half := m.visibleListEntries() / 2
		if half < 1 {
			half = 1
		}
		m.listCursor += half
		if m.listCursor >= len(m.currentEntries) {
			m.listCursor = len(m.currentEntries) - 1
		}
		m.adjustListScroll()
	case "ctrl+u", "pgup":
		half := m.visibleListEntries() / 2
		if half < 1 {
			half = 1
		}
		m.listCursor -= half
		if m.listCursor < 0 {
			m.listCursor = 0
		}
		m.adjustListScroll()
	case "g", "home":
		m.listCursor = 0
		m.adjustListScroll()
	case "G", "end":
		m.listCursor = len(m.currentEntries) - 1
		m.adjustListScroll()
	case "enter":
		if m.listCursor < len(m.currentEntries) {
			return m, openURL(m.currentEntries[m.listCursor].URL)
		}
	case "left", "h":
		m.activePanel = panelTree
	}
	return m, nil
}

func (m *Model) updateCurrentEntries() {
	if len(m.flatTree) == 0 {
		m.currentEntries = nil
		return
	}
	if m.treeCursor >= len(m.flatTree) {
		m.treeCursor = len(m.flatTree) - 1
	}
	node := m.flatTree[m.treeCursor].Node
	m.currentEntries = node.AllEntries()
	m.listCursor = 0
	m.listOffset = 0
}

func (m Model) visibleListEntries() int {
	v := m.listPanelHeight() / entryHeight
	if v < 1 {
		return 1
	}
	return v
}

func (m *Model) adjustListScroll() {
	visible := m.visibleListEntries()
	off := scrollOff
	if off > visible/2 {
		off = visible / 2
	}
	if m.listCursor < m.listOffset+off {
		m.listOffset = m.listCursor - off
	}
	if m.listCursor >= m.listOffset+visible-off {
		m.listOffset = m.listCursor - visible + off + 1
	}
	if m.listOffset < 0 {
		m.listOffset = 0
	}
}

func (m Model) listPanelHeight() int {
	// height minus header, footer, borders
	h := m.height - 4
	if h < 1 {
		h = 1
	}
	return h
}

// View renders the UI.
func (m Model) View() tea.View {
	if m.width == 0 || m.height == 0 {
		return tea.NewView("Loading...")
	}

	treeWidth := m.width*3/10 - 2        // 30% minus borders
	listWidth := m.width - treeWidth - 6 // remaining minus borders/gaps
	contentHeight := m.height - 3        // minus footer

	if treeWidth < 10 {
		treeWidth = 10
	}
	if listWidth < 20 {
		listWidth = 20
	}
	if contentHeight < 3 {
		contentHeight = 3
	}

	tree := m.renderTree(treeWidth, contentHeight)
	list := m.renderList(listWidth, contentHeight)

	// Apply border styles
	treeBorder := inactiveBorderStyle
	listBorder := inactiveBorderStyle
	if m.activePanel == panelTree {
		treeBorder = activeBorderStyle
	} else {
		listBorder = activeBorderStyle
	}

	treePanel := treeBorder.Width(treeWidth).Height(contentHeight).Render(tree)
	listPanel := listBorder.Width(listWidth).Height(contentHeight).Render(list)

	body := lipgloss.JoinHorizontal(lipgloss.Top, treePanel, listPanel)

	footer := m.renderFooter()

	content := lipgloss.JoinVertical(lipgloss.Left, body, footer)

	v := tea.NewView(content)
	v.AltScreen = true
	return v
}

func (m Model) renderTree(width, height int) string {
	var b strings.Builder

	title := headerStyle.Render("Categories")
	b.WriteString(title)
	b.WriteString("\n\n")

	linesUsed := 2
	end := m.treeOffset + height - 2
	if end > len(m.flatTree) {
		end = len(m.flatTree)
	}
	for i := m.treeOffset; i < end; i++ {
		fn := m.flatTree[i]
		if linesUsed >= height {
			break
		}

		indent := strings.Repeat("  ", fn.Depth)
		icon := "  "
		if fn.Node.HasChildren() {
			if fn.Node.Expanded {
				icon = "▼ "
			} else {
				icon = "▶ "
			}
		}

		count := fn.Node.TotalEntries()
		label := fmt.Sprintf("%s%s%s (%d)", indent, icon, fn.Node.Name, count)

		// Truncate to width
		if len(label) > width {
			label = label[:width-1] + "…"
		}

		if i == m.treeCursor {
			label = treeSelectedStyle.Render(label)
		} else {
			label = treeNormalStyle.Render(label)
		}

		b.WriteString(label)
		b.WriteString("\n")
		linesUsed++
	}

	return b.String()
}

func (m Model) renderList(width, height int) string {
	var b strings.Builder

	// Title
	title := "Resources"
	if m.filtering && m.filterText != "" {
		title = fmt.Sprintf("Resources (filter: %s)", m.filterText)
	}
	b.WriteString(headerStyle.Render(title))
	b.WriteString("\n\n")

	if len(m.currentEntries) == 0 {
		b.WriteString(entryDescStyle.Render("  No entries"))
		return b.String()
	}

	linesUsed := 2

	visible := (height - 2) / entryHeight
	if visible < 1 {
		visible = 1
	}

	start := m.listOffset
	end := start + visible
	if end > len(m.currentEntries) {
		end = len(m.currentEntries)
	}

	for idx := start; idx < end; idx++ {
		if linesUsed+entryHeight > height {
			break
		}

		e := m.currentEntries[idx]
		selected := idx == m.listCursor

		// Use a safe width that accounts for Unicode characters (★, ⑂)
		// that some terminals render as 2 columns but lipgloss counts as 1.
		safeWidth := width - 2

		// Line 1: name + stars + forks
		stats := fmt.Sprintf("★ %d", e.Stars)
		if e.Forks > 0 {
			stats += fmt.Sprintf("  ⑂ %d", e.Forks)
		}
		name := e.Name
		statsW := lipgloss.Width(stats)
		maxName := safeWidth - statsW - 2 // 2 for minimum gap
		if maxName < 4 {
			maxName = 4
		}
		if lipgloss.Width(name) > maxName {
			name = truncateToWidth(name, maxName-1) + "…"
		}
		nameStr := entryNameStyle.Render(name)
		statsStr := entryDescStyle.Render(stats)
		padding := safeWidth - lipgloss.Width(nameStr) - lipgloss.Width(statsStr)
		if padding < 1 {
			padding = 1
		}
		line1 := nameStr + strings.Repeat(" ", padding) + statsStr

		// Line 2: URL
		url := e.URL
		if lipgloss.Width(url) > safeWidth {
			url = truncateToWidth(url, safeWidth-1) + "…"
		}
		line2 := entryURLStyle.Render(url)

		// Line 3: description
		desc := e.Description
		if lipgloss.Width(desc) > safeWidth {
			desc = truncateToWidth(desc, safeWidth-3) + "..."
		}
		line3 := entryDescStyle.Render(desc)

		// Line 4: status + last push
		statusStr := statusStyle(e.Status).Render(e.Status)
		lastPush := ""
		if !e.LastPush.IsZero() {
			lastPush = fmt.Sprintf("  Last push: %s", e.LastPush.Format("2006-01-02"))
		}
		line4 := statusStr + entryDescStyle.Render(lastPush)

		// Line 5: separator
		sepWidth := safeWidth
		if sepWidth < 1 {
			sepWidth = 1
		}
		line5 := entryDescStyle.Render(strings.Repeat("─", sepWidth))

		entry := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", line1, line2, line3, line4, line5)

		if selected && m.activePanel == panelList {
			entry = entrySelectedStyle.Render(entry)
		}

		b.WriteString(entry)
		b.WriteString("\n")
		linesUsed += entryHeight
	}

	// Scroll indicator
	if len(m.currentEntries) > visible {
		indicator := fmt.Sprintf(" %d-%d of %d", start+1, end, len(m.currentEntries))
		b.WriteString(footerStyle.Render(indicator))
	}

	return b.String()
}

func (m Model) renderFooter() string {
	if m.filtering {
		return filterPromptStyle.Render("/") + entryDescStyle.Render(m.filterText+"█")
	}
	help := " Tab:switch  j/k:nav  PgDn/PgUp:page  g/G:top/bottom  Enter:expand/open  /:filter  q:quit"
	return footerStyle.Render(help)
}

// openURLMsg is sent after attempting to open a URL.
type openURLMsg struct{ err error }

func openURL(url string) tea.Cmd {
	return func() tea.Msg {
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "darwin":
			cmd = exec.Command("open", url)
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", url)
		default:
			cmd = exec.Command("xdg-open", url)
		}
		return openURLMsg{err: cmd.Run()}
	}
}

// truncateToWidth truncates s to at most maxWidth visible columns.
func truncateToWidth(s string, maxWidth int) string {
	if maxWidth <= 0 {
		return ""
	}
	w := 0
	for i, r := range s {
		rw := lipgloss.Width(string(r))
		if w+rw > maxWidth {
			return s[:i]
		}
		w += rw
	}
	return s
}
