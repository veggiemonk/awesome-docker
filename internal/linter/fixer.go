package linter

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/veggiemonk/awesome-docker/internal/parser"
)

// attributionRe matches trailing author attributions like:
//
//	by [@author](url), by [@author][ref], by @author
//
// Also handles "Created by", "Maintained by" etc.
var attributionRe = regexp.MustCompile(`\s+(?:(?:[Cc]reated|[Mm]aintained|[Bb]uilt)\s+)?by\s+\[@[^\]]+\](?:\([^)]*\)|\[[^\]]*\])\.?$`)

// bareAttributionRe matches: by @author at end of line (no link).
var bareAttributionRe = regexp.MustCompile(`\s+by\s+@\w+\.?$`)

// sectionHeadingRe matches markdown headings.
var sectionHeadingRe = regexp.MustCompile(`^(#{1,6})\s+(.+?)(?:\s*<!--.*-->)?$`)

// RemoveAttribution strips author attribution from a description string.
func RemoveAttribution(desc string) string {
	desc = attributionRe.ReplaceAllString(desc, "")
	desc = bareAttributionRe.ReplaceAllString(desc, "")
	return strings.TrimSpace(desc)
}

// FormatEntry reconstructs a markdown list line from a parsed Entry.
func FormatEntry(e parser.Entry) string {
	desc := e.Description
	var markers []string
	for _, m := range e.Markers {
		switch m {
		case parser.MarkerAbandoned:
			markers = append(markers, ":skull:")
		case parser.MarkerPaid:
			markers = append(markers, ":yen:")
		case parser.MarkerWIP:
			markers = append(markers, ":construction:")
		case parser.MarkerStale:
			markers = append(markers, ":ice_cube:")
		}
	}
	if len(markers) > 0 {
		desc = strings.Join(markers, " ") + " " + desc
	}
	return fmt.Sprintf("- [%s](%s) - %s", e.Name, e.URL, desc)
}

// FixFile reads the README, fixes entries (capitalize, period, remove attribution,
// sort), and writes the result back.
func FixFile(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	fixCount := 0

	var headingLines []int
	for i, line := range lines {
		if sectionHeadingRe.MatchString(line) {
			headingLines = append(headingLines, i)
		}
	}

	// Process each heading block independently to match linter sort scope.
	for i, headingIdx := range headingLines {
		start := headingIdx + 1
		end := len(lines)
		if i+1 < len(headingLines) {
			end = headingLines[i+1]
		}

		var entryPositions []int
		var entries []parser.Entry
		for lineIdx := start; lineIdx < end; lineIdx++ {
			entry, err := parser.ParseEntry(lines[lineIdx], lineIdx+1)
			if err != nil {
				continue
			}
			entryPositions = append(entryPositions, lineIdx)
			entries = append(entries, entry)
		}
		if len(entries) == 0 {
			continue
		}

		var fixed []parser.Entry
		for _, e := range entries {
			f := FixEntry(e)
			f.Description = RemoveAttribution(f.Description)
			// Re-apply period after removing attribution (it may have been stripped)
			if len(f.Description) > 0 && !strings.HasSuffix(f.Description, ".") {
				f.Description += "."
			}
			fixed = append(fixed, f)
		}

		sorted := SortEntries(fixed)
		for j, e := range sorted {
			newLine := FormatEntry(e)
			lineIdx := entryPositions[j]
			if lines[lineIdx] != newLine {
				fixCount++
				lines[lineIdx] = newLine
			}
		}
	}

	if fixCount == 0 {
		return 0, nil
	}

	// Write back
	out, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	w := bufio.NewWriter(out)
	for i, line := range lines {
		w.WriteString(line)
		if i < len(lines)-1 {
			w.WriteString("\n")
		}
	}
	// Preserve trailing newline if original had one
	w.WriteString("\n")
	return fixCount, w.Flush()
}
