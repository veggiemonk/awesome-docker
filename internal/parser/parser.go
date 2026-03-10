package parser

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// entryRe matches: - [Name](URL) - Description
// Also handles optional markers/text between URL and " - " separator, e.g.:
//
//   - [Name](URL) :skull: - Description
//   - [Name](URL) (2) :skull: - Description
var entryRe = regexp.MustCompile(`^[-*]\s+\[([^\]]+)\]\(([^)]+)\)(.*?)\s+-\s+(.+)$`)

// headingRe matches markdown headings: # Title, ## Title, etc.
var headingRe = regexp.MustCompile(`^(#{1,6})\s+(.+?)(?:\s*<!--.*-->)?$`)

var markerDefs = []struct {
	text   string
	marker Marker
}{
	{text: ":skull:", marker: MarkerAbandoned},
	{text: ":yen:", marker: MarkerPaid},
	{text: ":construction:", marker: MarkerWIP},
	{text: ":ice_cube:", marker: MarkerStale},
}

// ParseEntry parses a single markdown list line into an Entry.
func ParseEntry(line string, lineNum int) (Entry, error) {
	m := entryRe.FindStringSubmatch(strings.TrimSpace(line))
	if m == nil {
		return Entry{}, fmt.Errorf("line %d: not a valid entry: %q", lineNum, line)
	}

	middle := m[3] // text between URL closing paren and " - "
	desc := m[4]
	var markers []Marker

	// Extract markers from both the middle section and the description
	for _, def := range markerDefs {
		if strings.Contains(middle, def.text) || strings.Contains(desc, def.text) {
			markers = append(markers, def.marker)
			middle = strings.ReplaceAll(middle, def.text, "")
			desc = strings.ReplaceAll(desc, def.text, "")
		}
	}
	desc = strings.TrimSpace(desc)

	return Entry{
		Name:        m[1],
		URL:         m[2],
		Description: desc,
		Markers:     markers,
		Line:        lineNum,
		Raw:         line,
	}, nil
}

// Parse reads a full README and returns a Document.
func Parse(r io.Reader) (Document, error) {
	scanner := bufio.NewScanner(r)
	var doc Document
	var allSections []struct {
		section Section
		level   int
	}

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// Check for heading
		if hm := headingRe.FindStringSubmatch(line); hm != nil {
			level := len(hm[1])
			title := strings.TrimSpace(hm[2])
			allSections = append(allSections, struct {
				section Section
				level   int
			}{
				section: Section{Title: title, Level: level, Line: lineNum},
				level:   level,
			})
			continue
		}

		// Check for entry (list item with link)
		if entry, err := ParseEntry(line, lineNum); err == nil {
			if len(allSections) > 0 {
				allSections[len(allSections)-1].section.Entries = append(
					allSections[len(allSections)-1].section.Entries, entry)
			}
			continue
		}

		// Everything else: preamble if no sections yet
		if len(allSections) == 0 {
			doc.Preamble = append(doc.Preamble, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return doc, err
	}

	// Build section tree by nesting based on heading level
	doc.Sections = buildTree(allSections)
	return doc, nil
}

func buildTree(flat []struct {
	section Section
	level   int
},
) []Section {
	if len(flat) == 0 {
		return nil
	}

	var result []Section
	for i := 0; i < len(flat); i++ {
		current := flat[i].section
		currentLevel := flat[i].level

		// Collect children: everything after this heading at a deeper level
		j := i + 1
		for j < len(flat) && flat[j].level > currentLevel {
			j++
		}
		if j > i+1 {
			current.Children = buildTree(flat[i+1 : j])
		}
		result = append(result, current)
		i = j - 1
	}
	return result
}
