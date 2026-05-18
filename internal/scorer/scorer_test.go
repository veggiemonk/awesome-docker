package scorer

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/veggiemonk/awesome-docker/internal/checker"
)

func TestScoreHealthy(t *testing.T) {
	info := checker.RepoInfo{
		PushedAt:   time.Now().AddDate(0, -3, 0),
		IsArchived: false,
		Stars:      100,
		HasLicense: true,
	}
	status := Score(info)
	if status != StatusHealthy {
		t.Errorf("status = %q, want %q", status, StatusHealthy)
	}
}

func TestScoreInactive(t *testing.T) {
	info := checker.RepoInfo{
		PushedAt:   time.Now().AddDate(-1, -6, 0),
		IsArchived: false,
	}
	status := Score(info)
	if status != StatusInactive {
		t.Errorf("status = %q, want %q", status, StatusInactive)
	}
}

func TestScoreStale(t *testing.T) {
	info := checker.RepoInfo{
		PushedAt:   time.Now().AddDate(-3, 0, 0),
		IsArchived: false,
	}
	status := Score(info)
	if status != StatusStale {
		t.Errorf("status = %q, want %q", status, StatusStale)
	}
}

func TestScoreArchived(t *testing.T) {
	info := checker.RepoInfo{
		PushedAt:   time.Now(),
		IsArchived: true,
	}
	status := Score(info)
	if status != StatusArchived {
		t.Errorf("status = %q, want %q", status, StatusArchived)
	}
}

func TestScoreDisabled(t *testing.T) {
	info := checker.RepoInfo{
		IsDisabled: true,
	}
	status := Score(info)
	if status != StatusDead {
		t.Errorf("status = %q, want %q", status, StatusDead)
	}
}

func TestGenerateReport(t *testing.T) {
	results := []ScoredEntry{
		{URL: "https://github.com/a/a", Name: "a/a", Status: StatusHealthy, Stars: 100, LastPush: time.Now()},
		{URL: "https://github.com/b/b", Name: "b/b", Status: StatusArchived, Stars: 50, LastPush: time.Now()},
		{
			URL:      "https://github.com/c/c",
			Name:     "c/c",
			Status:   StatusStale,
			Stars:    10,
			LastPush: time.Now().AddDate(-3, 0, 0),
		},
	}
	report := GenerateReport(results)
	if !strings.Contains(report, "Healthy: 1") {
		t.Error("report should contain 'Healthy: 1'")
	}
	if !strings.Contains(report, "Archived: 1") {
		t.Error("report should contain 'Archived: 1'")
	}
	if !strings.Contains(report, "Stale") {
		t.Error("report should contain 'Stale'")
	}
}

func TestGenerateReportShowsAllEntries(t *testing.T) {
	const entryCount = 55
	results := make([]ScoredEntry, 0, entryCount)
	for i := range entryCount {
		results = append(results, ScoredEntry{
			URL:      fmt.Sprintf("https://github.com/stale/%d", i),
			Name:     fmt.Sprintf("stale/%d", i),
			Status:   StatusStale,
			Stars:    i,
			LastPush: time.Now().AddDate(-3, 0, 0),
		})
	}

	report := GenerateReport(results)
	if strings.Contains(report, "... and") {
		t.Fatal("report should not be truncated")
	}
	if !strings.Contains(report, fmt.Sprintf("stale/%d", entryCount-1)) {
		t.Fatal("report should contain all entries")
	}
}

func TestGenerateJSONReport(t *testing.T) {
	results := []ScoredEntry{
		{
			URL:      "https://github.com/a/a",
			Name:     "a/a",
			Status:   StatusHealthy,
			Stars:    100,
			LastPush: time.Now(),
		},
		{
			URL:      "https://github.com/b/b",
			Name:     "b/b",
			Status:   StatusStale,
			Stars:    50,
			LastPush: time.Now().AddDate(-3, 0, 0),
		},
	}

	data, err := GenerateJSONReport(results)
	if err != nil {
		t.Fatalf("GenerateJSONReport() error = %v", err)
	}

	var report ReportData
	if err := json.Unmarshal(data, &report); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if report.Total != 2 {
		t.Fatalf("report.Total = %d, want 2", report.Total)
	}
	if report.Summary.Healthy != 1 || report.Summary.Stale != 1 {
		t.Fatalf("summary = %+v, want healthy=1 stale=1", report.Summary)
	}
	if len(report.Entries) != 2 {
		t.Fatalf("len(report.Entries) = %d, want 2", len(report.Entries))
	}
	if len(report.ByStatus[StatusStale]) != 1 {
		t.Fatalf("len(report.ByStatus[stale]) = %d, want 1", len(report.ByStatus[StatusStale]))
	}
}

func TestScoreAll(t *testing.T) {
	infos := []checker.RepoInfo{
		{Owner: "a", Name: "a", PushedAt: time.Now(), Stars: 10},
		{Owner: "b", Name: "b", PushedAt: time.Now().AddDate(-3, 0, 0), Stars: 5},
	}
	scored := ScoreAll(infos)
	if len(scored) != 2 {
		t.Fatalf("scored = %d, want 2", len(scored))
	}
	if scored[0].Status != StatusHealthy {
		t.Errorf("first = %q, want healthy", scored[0].Status)
	}
	if scored[1].Status != StatusStale {
		t.Errorf("second = %q, want stale", scored[1].Status)
	}
}
