package tui

import "charm.land/lipgloss/v2"

var (
	// Panel borders
	activeBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#7D56F4"))

	inactiveBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#555555"))

	// Tree styles
	treeSelectedStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF79C6")).Background(lipgloss.Color("#3B2D50"))
	treeNormalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC"))

	// Entry styles
	entryNameStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#50FA7B"))
	entryURLStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Italic(true)
	entryDescStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC"))

	// Status badge styles
	statusHealthyStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#50FA7B")).Bold(true)
	statusInactiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB86C"))
	statusStaleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#F1FA8C"))
	statusArchivedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5555")).Bold(true)
	statusDeadStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666")).Strikethrough(true)

	// Selected entry
	entrySelectedStyle = lipgloss.NewStyle().Background(lipgloss.Color("#44475A"))

	// Header
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#BD93F9")).Padding(0, 1)

	// Footer
	footerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))

	// Filter
	filterPromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF79C6")).Bold(true)
)

func statusStyle(status string) lipgloss.Style {
	switch status {
	case "healthy":
		return statusHealthyStyle
	case "inactive":
		return statusInactiveStyle
	case "stale":
		return statusStaleStyle
	case "archived":
		return statusArchivedStyle
	case "dead":
		return statusDeadStyle
	default:
		return lipgloss.NewStyle()
	}
}
