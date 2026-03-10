package parser

import (
	"os"
	"strings"
	"testing"
)

func TestParseEntry(t *testing.T) {
	line := `- [Docker Desktop](https://www.docker.com/products/docker-desktop/) - Official native app. Only for Windows and MacOS.`
	entry, err := ParseEntry(line, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if entry.Name != "Docker Desktop" {
		t.Errorf("name = %q, want %q", entry.Name, "Docker Desktop")
	}
	if entry.URL != "https://www.docker.com/products/docker-desktop/" {
		t.Errorf("url = %q, want %q", entry.URL, "https://www.docker.com/products/docker-desktop/")
	}
	if entry.Description != "Official native app. Only for Windows and MacOS." {
		t.Errorf("description = %q, want %q", entry.Description, "Official native app. Only for Windows and MacOS.")
	}
	if len(entry.Markers) != 0 {
		t.Errorf("markers = %v, want empty", entry.Markers)
	}
}

func TestParseEntryWithMarkers(t *testing.T) {
	line := `- [Docker Swarm](https://github.com/docker/swarm) - Swarm clustering system. :skull:`
	entry, err := ParseEntry(line, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if entry.Name != "Docker Swarm" {
		t.Errorf("name = %q, want %q", entry.Name, "Docker Swarm")
	}
	if len(entry.Markers) != 1 || entry.Markers[0] != MarkerAbandoned {
		t.Errorf("markers = %v, want [MarkerAbandoned]", entry.Markers)
	}
	if strings.Contains(entry.Description, ":skull:") {
		t.Errorf("description should not contain marker text, got %q", entry.Description)
	}
}

func TestParseEntryMultipleMarkers(t *testing.T) {
	line := `- [SomeProject](https://example.com) - A project. :yen: :construction:`
	entry, err := ParseEntry(line, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(entry.Markers) != 2 {
		t.Fatalf("markers count = %d, want 2", len(entry.Markers))
	}
}

func TestParseEntryMarkersCanonicalOrder(t *testing.T) {
	line := `- [SomeProject](https://example.com) - :construction: A project. :skull:`
	entry, err := ParseEntry(line, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(entry.Markers) != 2 {
		t.Fatalf("markers count = %d, want 2", len(entry.Markers))
	}
	if entry.Markers[0] != MarkerAbandoned || entry.Markers[1] != MarkerWIP {
		t.Fatalf("marker order = %v, want [MarkerAbandoned MarkerWIP]", entry.Markers)
	}
}

func TestParseDocument(t *testing.T) {
	input := `# Awesome Docker

> A curated list

# Contents

- [Projects](#projects)

# Legend

- Abandoned :skull:

# Projects

## Tools

- [ToolA](https://github.com/a/a) - Does A.
- [ToolB](https://github.com/b/b) - Does B. :skull:

## Services

- [ServiceC](https://example.com/c) - Does C. :yen:
`
	doc, err := Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(doc.Sections) == 0 {
		t.Fatal("expected at least one section")
	}
	// Find the "Projects" section
	var projects *Section
	for i := range doc.Sections {
		if doc.Sections[i].Title == "Projects" {
			projects = &doc.Sections[i]
			break
		}
	}
	if projects == nil {
		t.Fatal("expected a Projects section")
	}
	if len(projects.Children) != 2 {
		t.Errorf("projects children = %d, want 2", len(projects.Children))
	}
	if projects.Children[0].Title != "Tools" {
		t.Errorf("first child = %q, want %q", projects.Children[0].Title, "Tools")
	}
	if len(projects.Children[0].Entries) != 2 {
		t.Errorf("Tools entries = %d, want 2", len(projects.Children[0].Entries))
	}
}

func TestParseNotAnEntry(t *testing.T) {
	_, err := ParseEntry("- Abandoned :skull:", 1)
	if err == nil {
		t.Error("expected error for non-entry list item")
	}
}

func TestParseRealREADME(t *testing.T) {
	f, err := os.Open("../../README.md")
	if err != nil {
		t.Skip("README.md not found, skipping integration test")
	}
	defer f.Close()

	doc, err := Parse(f)
	if err != nil {
		t.Fatalf("failed to parse README: %v", err)
	}

	if len(doc.Sections) == 0 {
		t.Error("expected sections")
	}

	total := countEntries(doc.Sections)
	if total < 100 {
		t.Errorf("expected at least 100 entries, got %d", total)
	}
	t.Logf("Parsed %d sections, %d total entries", len(doc.Sections), total)
}

func countEntries(sections []Section) int {
	n := 0
	for _, s := range sections {
		n += len(s.Entries)
		n += countEntries(s.Children)
	}
	return n
}
