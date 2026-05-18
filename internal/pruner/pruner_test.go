package pruner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/veggiemonk/awesome-docker/internal/cache"
)

func TestTargetURLs(t *testing.T) {
	hc := &cache.HealthCache{Entries: []cache.HealthEntry{
		{URL: "https://github.com/A/x", Status: "archived"},
		{URL: "https://github.com/B/y", Status: "stale"},
		{URL: "https://github.com/C/z", Status: "healthy"},
		{URL: "https://github.com/D/w", Status: "inactive"},
	}}
	got := TargetURLs(hc, []string{"archived", "stale"})
	if len(got) != 2 {
		t.Fatalf("want 2 targets, got %d", len(got))
	}
	if _, ok := got["https://github.com/a/x"]; !ok {
		t.Errorf("expected lowercased URL key for archived entry")
	}
}

func TestTargetsFromReport(t *testing.T) {
	r := strings.NewReader(`# Health Report

## Summary

- Stale (2+ years): 2

## Archived (should mark :skull:)

- [a/keep](https://github.com/A/Keep) - Stars: 1 - Last push: 2024-01-01

## Stale (2+ years inactive)

- [b/drop](https://github.com/b/drop) - Stars: 2 - Last push: 2020-01-01

## Inactive (1-2 years)

- [c/skip](https://github.com/c/skip) - Stars: 3 - Last push: 2025-01-01
`)
	targets, err := TargetsFromReport(r, []string{"archived", "stale"})
	if err != nil {
		t.Fatal(err)
	}
	if len(targets) != 2 {
		t.Fatalf("want 2, got %d: %v", len(targets), targets)
	}
	if _, ok := targets["https://github.com/a/keep"]; !ok {
		t.Errorf("missing archived entry (case-insensitive)")
	}
	if _, ok := targets["https://github.com/b/drop"]; !ok {
		t.Errorf("missing stale entry")
	}
	if _, ok := targets["https://github.com/c/skip"]; ok {
		t.Errorf("inactive entry should not have been picked up")
	}
}

func TestPruneREADME(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "README.md")
	content := `# Header

## Tools

- [keep](https://github.com/keep/me) - Healthy project.
- [drop](https://github.com/drop/me) - Stale project.
- [also-keep](https://github.com/also/keep) - Another one.
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	targets := map[string]cache.HealthEntry{
		"https://github.com/drop/me": {URL: "https://github.com/drop/me", Status: "stale"},
	}
	res, err := PruneREADME(path, targets, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Removed) != 1 {
		t.Fatalf("want 1 removed, got %d", len(res.Removed))
	}
	if res.Removed[0].URL != "https://github.com/drop/me" {
		t.Errorf("unexpected removed URL: %s", res.Removed[0].URL)
	}

	out, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(out), "drop/me") {
		t.Errorf("expected drop/me to be removed from README, got:\n%s", out)
	}
	if !strings.Contains(string(out), "keep/me") || !strings.Contains(string(out), "also/keep") {
		t.Errorf("expected other entries to be preserved, got:\n%s", out)
	}
}

func TestPruneREADMEDryRun(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "README.md")
	content := "## X\n\n- [drop](https://github.com/drop/me) - Stale.\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	targets := map[string]cache.HealthEntry{
		"https://github.com/drop/me": {URL: "https://github.com/drop/me", Status: "stale"},
	}
	res, err := PruneREADME(path, targets, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Removed) != 1 {
		t.Fatalf("want 1 removed (preview), got %d", len(res.Removed))
	}
	got, _ := os.ReadFile(path)
	if string(got) != content {
		t.Errorf("dry-run modified file: %q", got)
	}
}

func TestPruneREADMENotFound(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "README.md")
	if err := os.WriteFile(path, []byte("## X\n\n- [k](https://github.com/k/v) - Keep.\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	targets := map[string]cache.HealthEntry{
		"https://github.com/gone/missing": {URL: "https://github.com/gone/missing", Status: "stale"},
	}
	res, err := PruneREADME(path, targets, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Removed) != 0 {
		t.Errorf("want 0 removed, got %d", len(res.Removed))
	}
	if len(res.NotFound) != 1 || res.NotFound[0] != "https://github.com/gone/missing" {
		t.Errorf("want gone/missing in NotFound, got %v", res.NotFound)
	}
}

func TestPruneCache(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cache.yaml")
	hc := &cache.HealthCache{Entries: []cache.HealthEntry{
		{URL: "https://github.com/a/keep", Status: "healthy"},
		{URL: "https://github.com/b/drop", Status: "stale"},
	}}
	if err := cache.SaveHealthCache(path, hc); err != nil {
		t.Fatal(err)
	}
	targets := map[string]cache.HealthEntry{
		"https://github.com/b/drop": {URL: "https://github.com/b/drop", Status: "stale"},
	}
	n, err := PruneCache(path, hc, targets, false)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Errorf("want 1 dropped, got %d", n)
	}
	if len(hc.Entries) != 1 || hc.Entries[0].URL != "https://github.com/a/keep" {
		t.Errorf("unexpected remaining entries: %v", hc.Entries)
	}
}
