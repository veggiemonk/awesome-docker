package tui

import (
	"testing"

	"github.com/veggiemonk/awesome-docker/internal/cache"
)

func TestBuildTree(t *testing.T) {
	entries := []cache.HealthEntry{
		{URL: "https://github.com/a/b", Name: "a/b", Category: "Projects > Networking", Description: "desc1"},
		{URL: "https://github.com/c/d", Name: "c/d", Category: "Projects > Networking", Description: "desc2"},
		{URL: "https://github.com/e/f", Name: "e/f", Category: "Projects > Security", Description: "desc3"},
		{URL: "https://github.com/g/h", Name: "g/h", Category: "Docker Images > Base Tools", Description: "desc4"},
		{URL: "https://github.com/i/j", Name: "i/j", Category: "", Description: "no category"},
	}

	roots := BuildTree(entries)

	// Should have 3 roots: Docker Images, Projects, Uncategorized (sorted)
	if len(roots) != 3 {
		t.Fatalf("expected 3 roots, got %d", len(roots))
	}

	if roots[0].Name != "Docker Images" {
		t.Errorf("expected first root 'Docker Images', got %q", roots[0].Name)
	}
	if roots[1].Name != "Projects" {
		t.Errorf("expected second root 'Projects', got %q", roots[1].Name)
	}
	if roots[2].Name != "Uncategorized" {
		t.Errorf("expected third root 'Uncategorized', got %q", roots[2].Name)
	}

	// Projects > Networking should have 2 entries
	projects := roots[1]
	if len(projects.Children) != 2 {
		t.Fatalf("expected 2 children under Projects, got %d", len(projects.Children))
	}
	networking := projects.Children[0] // Networking < Security alphabetically
	if networking.Name != "Networking" {
		t.Errorf("expected 'Networking', got %q", networking.Name)
	}
	if len(networking.Entries) != 2 {
		t.Errorf("expected 2 entries in Networking, got %d", len(networking.Entries))
	}
}

func TestBuildTreeEmpty(t *testing.T) {
	roots := BuildTree(nil)
	if len(roots) != 0 {
		t.Errorf("expected 0 roots for nil input, got %d", len(roots))
	}
}

func TestTotalEntries(t *testing.T) {
	entries := []cache.HealthEntry{
		{URL: "https://a", Category: "A > B"},
		{URL: "https://b", Category: "A > B"},
		{URL: "https://c", Category: "A > C"},
		{URL: "https://d", Category: "A"},
	}
	roots := BuildTree(entries)
	if len(roots) != 1 {
		t.Fatalf("expected 1 root, got %d", len(roots))
	}
	if roots[0].TotalEntries() != 4 {
		t.Errorf("expected 4 total entries, got %d", roots[0].TotalEntries())
	}
}

func TestFlattenVisible(t *testing.T) {
	entries := []cache.HealthEntry{
		{URL: "https://a", Category: "A > B"},
		{URL: "https://b", Category: "A > C"},
	}
	roots := BuildTree(entries)

	// Initially not expanded, should see just root
	flat := FlattenVisible(roots)
	if len(flat) != 1 {
		t.Fatalf("expected 1 visible node (collapsed), got %d", len(flat))
	}
	if flat[0].Depth != 0 {
		t.Errorf("expected depth 0, got %d", flat[0].Depth)
	}

	// Expand root
	roots[0].Expanded = true
	flat = FlattenVisible(roots)
	if len(flat) != 3 {
		t.Fatalf("expected 3 visible nodes (expanded), got %d", len(flat))
	}
	if flat[1].Depth != 1 {
		t.Errorf("expected depth 1 for child, got %d", flat[1].Depth)
	}
}

func TestAllEntries(t *testing.T) {
	entries := []cache.HealthEntry{
		{URL: "https://a", Category: "A > B"},
		{URL: "https://b", Category: "A"},
	}
	roots := BuildTree(entries)
	all := roots[0].AllEntries()
	if len(all) != 2 {
		t.Errorf("expected 2 entries from AllEntries, got %d", len(all))
	}
}
