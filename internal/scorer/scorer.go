package scorer

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/veggiemonk/awesome-docker/internal/cache"
	"github.com/veggiemonk/awesome-docker/internal/checker"
)

// Status represents the health status of an entry.
type Status string

const (
	StatusHealthy  Status = "healthy"
	StatusInactive Status = "inactive" // 1-2 years since last push
	StatusStale    Status = "stale"    // 2+ years since last push
	StatusArchived Status = "archived"
	StatusDead     Status = "dead" // disabled or 404
)

// ScoredEntry is a repo with its computed health status.
type ScoredEntry struct {
	URL         string
	Name        string
	Status      Status
	Stars       int
	Forks       int
	HasLicense  bool
	LastPush    time.Time
	Category    string
	Description string
}

// ReportSummary contains grouped status counts.
type ReportSummary struct {
	Healthy  int `json:"healthy"`
	Inactive int `json:"inactive"`
	Stale    int `json:"stale"`
	Archived int `json:"archived"`
	Dead     int `json:"dead"`
}

// ReportData is the full machine-readable report model.
type ReportData struct {
	GeneratedAt time.Time                `json:"generated_at"`
	Total       int                      `json:"total"`
	Summary     ReportSummary            `json:"summary"`
	Entries     []ScoredEntry            `json:"entries"`
	ByStatus    map[Status][]ScoredEntry `json:"by_status"`
}

// Score computes the health status of a GitHub repo.
func Score(info checker.RepoInfo) Status {
	if info.IsDisabled {
		return StatusDead
	}
	if info.IsArchived {
		return StatusArchived
	}

	twoYearsAgo := time.Now().AddDate(-2, 0, 0)
	oneYearAgo := time.Now().AddDate(-1, 0, 0)

	if info.PushedAt.Before(twoYearsAgo) {
		return StatusStale
	}
	if info.PushedAt.Before(oneYearAgo) {
		return StatusInactive
	}
	return StatusHealthy
}

// ScoreAll scores a batch of repo infos.
func ScoreAll(infos []checker.RepoInfo) []ScoredEntry {
	results := make([]ScoredEntry, len(infos))
	for i, info := range infos {
		results[i] = ScoredEntry{
			URL:        info.URL,
			Name:       fmt.Sprintf("%s/%s", info.Owner, info.Name),
			Status:     Score(info),
			Stars:      info.Stars,
			Forks:      info.Forks,
			HasLicense: info.HasLicense,
			LastPush:   info.PushedAt,
		}
	}
	return results
}

// ToCacheEntries converts scored entries to cache format.
func ToCacheEntries(scored []ScoredEntry) []cache.HealthEntry {
	entries := make([]cache.HealthEntry, len(scored))
	now := time.Now().UTC()
	for i, s := range scored {
		entries[i] = cache.HealthEntry{
			URL:         s.URL,
			Name:        s.Name,
			Status:      string(s.Status),
			Stars:       s.Stars,
			Forks:       s.Forks,
			HasLicense:  s.HasLicense,
			LastPush:    s.LastPush,
			CheckedAt:   now,
			Category:    s.Category,
			Description: s.Description,
		}
	}
	return entries
}

// GenerateReport produces a Markdown health report.
func GenerateReport(scored []ScoredEntry) string {
	var b strings.Builder

	data := BuildReportData(scored)
	groups := data.ByStatus

	fmt.Fprintf(&b, "# Health Report\n\n")
	fmt.Fprintf(&b, "**Generated:** %s\n\n", data.GeneratedAt.Format(time.RFC3339))
	fmt.Fprintf(&b, "**Total:** %d repositories\n\n", data.Total)

	fmt.Fprintf(&b, "## Summary\n\n")
	fmt.Fprintf(&b, "- Healthy: %d\n", data.Summary.Healthy)
	fmt.Fprintf(&b, "- Inactive (1-2 years): %d\n", data.Summary.Inactive)
	fmt.Fprintf(&b, "- Stale (2+ years): %d\n", data.Summary.Stale)
	fmt.Fprintf(&b, "- Archived: %d\n", data.Summary.Archived)
	fmt.Fprintf(&b, "- Dead: %d\n\n", data.Summary.Dead)

	writeSection := func(title string, status Status) {
		entries := groups[status]
		if len(entries) == 0 {
			return
		}
		fmt.Fprintf(&b, "## %s\n\n", title)
		for _, e := range entries {
			fmt.Fprintf(&b, "- [%s](%s) - Stars: %d - Last push: %s\n",
				e.Name, e.URL, e.Stars, e.LastPush.Format("2006-01-02"))
		}
		b.WriteString("\n")
	}

	writeSection("Archived (should mark :skull:)", StatusArchived)
	writeSection("Stale (2+ years inactive)", StatusStale)
	writeSection("Inactive (1-2 years)", StatusInactive)

	return b.String()
}

// BuildReportData returns full report data for machine-readable and markdown rendering.
func BuildReportData(scored []ScoredEntry) ReportData {
	groups := map[Status][]ScoredEntry{}
	for _, s := range scored {
		groups[s.Status] = append(groups[s.Status], s)
	}

	return ReportData{
		GeneratedAt: time.Now().UTC(),
		Total:       len(scored),
		Summary: ReportSummary{
			Healthy:  len(groups[StatusHealthy]),
			Inactive: len(groups[StatusInactive]),
			Stale:    len(groups[StatusStale]),
			Archived: len(groups[StatusArchived]),
			Dead:     len(groups[StatusDead]),
		},
		Entries:  scored,
		ByStatus: groups,
	}
}

// GenerateJSONReport returns the full report as pretty-printed JSON.
func GenerateJSONReport(scored []ScoredEntry) ([]byte, error) {
	data := BuildReportData(scored)
	return json.MarshalIndent(data, "", "  ")
}
