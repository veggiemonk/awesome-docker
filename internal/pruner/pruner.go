// Package pruner owns the removal of README entries by health status.
//
// Why it exists: maintenance regularly produces a list of archived/stale
// projects (see scorer + cache). Pruner is the seam that translates that list
// into a concrete edit of README.md and config/health_cache.yaml, so the README
// stays in lockstep with the cache instead of drifting via ad-hoc edits.
package pruner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/veggiemonk/awesome-docker/internal/cache"
	"github.com/veggiemonk/awesome-docker/internal/parser"
)

// Removed describes a single entry removed from the README.
type Removed struct {
	URL    string
	Name   string
	Status string
	Line   int
}

// Result summarizes a prune run.
type Result struct {
	Removed []Removed
	// URLs in the target set that didn't appear in the README (already gone,
	// non-GitHub indirection, or URL drift between cache and README).
	NotFound []string
}

// TargetURLs returns the URL set selected by the given statuses from the cache.
func TargetURLs(hc *cache.HealthCache, statuses []string) map[string]cache.HealthEntry {
	want := make(map[string]bool, len(statuses))
	for _, s := range statuses {
		want[strings.TrimSpace(strings.ToLower(s))] = true
	}
	out := make(map[string]cache.HealthEntry)
	for i := range hc.Entries {
		e := &hc.Entries[i]
		if want[strings.ToLower(e.Status)] {
			out[normalizeURL(e.URL)] = *e
		}
	}
	return out
}

// PruneREADME removes lines whose entry URL is in targets and writes the
// result back to path. If dryRun is true, the file is not modified.
func PruneREADME(path string, targets map[string]cache.HealthEntry, dryRun bool) (Result, error) {
	f, err := os.Open(path) //nolint:gosec
	if err != nil {
		return Result{}, fmt.Errorf("open %s: %w", path, err)
	}
	lines, err := readLines(f)
	f.Close()
	if err != nil {
		return Result{}, fmt.Errorf("read %s: %w", path, err)
	}

	var (
		kept    = make([]string, 0, len(lines))
		removed []Removed
		hit     = make(map[string]bool, len(targets))
	)

	for i, line := range lines {
		entry, perr := parser.ParseEntry(line, i+1)
		if perr != nil {
			kept = append(kept, line)
			continue
		}
		key := normalizeURL(entry.URL)
		meta, ok := targets[key]
		if !ok {
			kept = append(kept, line)
			continue
		}
		hit[key] = true
		removed = append(removed, Removed{
			URL:    entry.URL,
			Name:   entry.Name,
			Status: meta.Status,
			Line:   i + 1,
		})
	}

	res := Result{Removed: removed}
	for k := range targets {
		if !hit[k] {
			res.NotFound = append(res.NotFound, targets[k].URL)
		}
	}
	sort.Strings(res.NotFound)

	if dryRun || len(removed) == 0 {
		return res, nil
	}

	if err := writeLines(path, kept); err != nil {
		return res, fmt.Errorf("write %s: %w", path, err)
	}
	return res, nil
}

// PruneCache drops entries whose normalized URL is in targets and writes the
// cache back to path. Safe to call when len(targets) == 0 (no-op).
func PruneCache(path string, hc *cache.HealthCache, targets map[string]cache.HealthEntry, dryRun bool) (int, error) {
	if len(targets) == 0 {
		return 0, nil
	}
	kept := hc.Entries[:0]
	for i := range hc.Entries {
		e := &hc.Entries[i]
		if _, drop := targets[normalizeURL(e.URL)]; drop {
			continue
		}
		kept = append(kept, *e)
	}
	dropped := len(hc.Entries) - len(kept)
	hc.Entries = kept
	if dryRun || dropped == 0 {
		return dropped, nil
	}
	if err := cache.SaveHealthCache(path, hc); err != nil {
		return dropped, err
	}
	return dropped, nil
}

// reportSectionRe matches markdown health-report section headings:
//
//	## Archived (should mark :skull:)
//	## Stale (2+ years inactive)
//	## Inactive (1-2 years)
var reportSectionRe = regexp.MustCompile(`(?i)^##\s+(archived|stale|inactive|dead|healthy)\b`)

// reportEntryRe matches: "- [name](url) - Stars: N - Last push: YYYY-MM-DD"
var reportEntryRe = regexp.MustCompile(`^-\s+\[([^\]]+)\]\((https?://[^)]+)\)`)

// TargetsFromReport parses a markdown health report (same format as the
// `report` subcommand emits) and returns the URL set whose section heading
// matches one of the given statuses.
func TargetsFromReport(r io.Reader, statuses []string) (map[string]cache.HealthEntry, error) {
	want := make(map[string]bool, len(statuses))
	for _, s := range statuses {
		want[strings.TrimSpace(strings.ToLower(s))] = true
	}
	out := make(map[string]cache.HealthEntry)
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	var current string
	for sc.Scan() {
		line := sc.Text()
		if m := reportSectionRe.FindStringSubmatch(line); m != nil {
			current = strings.ToLower(m[1])
			continue
		}
		if !want[current] {
			continue
		}
		if m := reportEntryRe.FindStringSubmatch(line); m != nil {
			url := strings.TrimSpace(m[2])
			out[normalizeURL(url)] = cache.HealthEntry{
				URL:    url,
				Name:   m[1],
				Status: current,
			}
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func normalizeURL(u string) string {
	u = strings.TrimSpace(u)
	u = strings.TrimSuffix(u, "/")
	u = strings.ToLower(u)
	return u
}

func readLines(r *os.File) ([]string, error) {
	var lines []string
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines, sc.Err()
}

func writeLines(path string, lines []string) error {
	out, err := os.Create(path) //nolint:gosec
	if err != nil {
		return err
	}
	defer out.Close()
	w := bufio.NewWriter(out)
	for i, line := range lines {
		if _, err := w.WriteString(line); err != nil {
			return err
		}
		if i < len(lines)-1 {
			if err := w.WriteByte('\n'); err != nil {
				return err
			}
		}
	}
	if err := w.WriteByte('\n'); err != nil {
		return err
	}
	return w.Flush()
}
