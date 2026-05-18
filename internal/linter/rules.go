package linter

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/veggiemonk/awesome-docker/internal/parser"
)

// Rule identifies a linting rule.
type Rule string

const (
	RuleDescriptionCapital Rule = "description-capital"
	RuleDescriptionPeriod  Rule = "description-period"
	RuleSorted             Rule = "sorted"
	RuleDuplicateURL       Rule = "duplicate-url"
)

// Severity of a lint issue.
type Severity int

const (
	SeverityError Severity = iota
	SeverityWarning
)

// Issue is a single lint problem found.
type Issue struct {
	Rule     Rule
	Message  string
	Severity Severity
	Line     int
}

func (i Issue) String() string {
	sev := "ERROR"
	if i.Severity == SeverityWarning {
		sev = "WARN"
	}
	return fmt.Sprintf("[%s] line %d: %s (%s)", sev, i.Line, i.Message, i.Rule)
}

// CheckEntry validates a single entry against formatting rules.
func CheckEntry(e parser.Entry) []Issue {
	var issues []Issue

	if first, ok := firstLetter(e.Description); ok && !unicode.IsUpper(first) {
		issues = append(issues, Issue{
			Rule:     RuleDescriptionCapital,
			Severity: SeverityError,
			Line:     e.Line,
			Message:  fmt.Sprintf("%q: description should start with a capital letter", e.Name),
		})
	}

	if len(e.Description) > 0 && !strings.HasSuffix(e.Description, ".") {
		issues = append(issues, Issue{
			Rule:     RuleDescriptionPeriod,
			Severity: SeverityError,
			Line:     e.Line,
			Message:  fmt.Sprintf("%q: description should end with a period", e.Name),
		})
	}

	return issues
}

// CheckSorted verifies entries are in alphabetical order (case-insensitive).
func CheckSorted(entries []parser.Entry) []Issue {
	var issues []Issue
	for i := 1; i < len(entries); i++ {
		prev := strings.ToLower(entries[i-1].Name)
		curr := strings.ToLower(entries[i].Name)
		if prev > curr {
			issues = append(issues, Issue{
				Rule:     RuleSorted,
				Severity: SeverityError,
				Line:     entries[i].Line,
				Message: fmt.Sprintf(
					"%q should come before %q (alphabetical order)",
					entries[i].Name,
					entries[i-1].Name,
				),
			})
		}
	}
	return issues
}

// CheckDuplicates finds entries with the same URL across the entire document.
func CheckDuplicates(entries []parser.Entry) []Issue {
	var issues []Issue
	seen := make(map[string]int) // URL -> first line number
	for _, e := range entries {
		url := strings.TrimRight(e.URL, "/")
		if firstLine, exists := seen[url]; exists {
			issues = append(issues, Issue{
				Rule:     RuleDuplicateURL,
				Severity: SeverityError,
				Line:     e.Line,
				Message:  fmt.Sprintf("duplicate URL %q (first seen at line %d)", e.URL, firstLine),
			})
		} else {
			seen[url] = e.Line
		}
	}
	return issues
}

// firstLetter returns the first unicode letter in s and true, or zero and false if none.
func firstLetter(s string) (rune, bool) {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return r, true
		}
	}
	return 0, false
}

// FixEntry returns a copy of the entry with auto-fixable issues corrected.
func FixEntry(e parser.Entry) parser.Entry {
	fixed := e
	if len(fixed.Description) > 0 {
		// Capitalize first letter (find it, may not be at index 0)
		runes := []rune(fixed.Description)
		for i, r := range runes {
			if unicode.IsLetter(r) {
				runes[i] = unicode.ToUpper(r)
				break
			}
		}
		fixed.Description = string(runes)

		// Ensure period at end
		if !strings.HasSuffix(fixed.Description, ".") {
			fixed.Description += "."
		}
	}
	return fixed
}

// SortEntries returns a sorted copy of entries (case-insensitive by Name).
func SortEntries(entries []parser.Entry) []parser.Entry {
	sorted := make([]parser.Entry, len(entries))
	copy(sorted, entries)
	sort.Slice(sorted, func(i, j int) bool {
		return strings.ToLower(sorted[i].Name) < strings.ToLower(sorted[j].Name)
	})
	return sorted
}
