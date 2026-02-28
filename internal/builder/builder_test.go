package builder

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuild(t *testing.T) {
	dir := t.TempDir()

	md := "# Test List\n\n- [Example](https://example.com) - A test entry.\n"
	mdPath := filepath.Join(dir, "README.md")
	if err := os.WriteFile(mdPath, []byte(md), 0o644); err != nil {
		t.Fatal(err)
	}

	tmpl := `<!DOCTYPE html>
<html>
<body>
<div id="md"></div>
</body>
</html>`
	tmplPath := filepath.Join(dir, "template.html")
	if err := os.WriteFile(tmplPath, []byte(tmpl), 0o644); err != nil {
		t.Fatal(err)
	}

	outPath := filepath.Join(dir, "index.html")
	if err := Build(mdPath, tmplPath, outPath); err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	content, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatal(err)
	}

	html := string(content)
	if !strings.Contains(html, "Test List") {
		t.Error("expected 'Test List' in output")
	}
	if !strings.Contains(html, "https://example.com") {
		t.Error("expected link in output")
	}
}

func TestBuildWithSectionPlaceholder(t *testing.T) {
	dir := t.TempDir()

	md := "# Hello\n\nWorld.\n"
	mdPath := filepath.Join(dir, "README.md")
	if err := os.WriteFile(mdPath, []byte(md), 0o644); err != nil {
		t.Fatal(err)
	}

	// This matches the actual template format
	tmpl := `<!DOCTYPE html>
<html>
<body>
<section id="md" class="main-content"></section>
</body>
</html>`
	tmplPath := filepath.Join(dir, "template.html")
	if err := os.WriteFile(tmplPath, []byte(tmpl), 0o644); err != nil {
		t.Fatal(err)
	}

	outPath := filepath.Join(dir, "index.html")
	if err := Build(mdPath, tmplPath, outPath); err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	content, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(content), "Hello") {
		t.Error("expected 'Hello' in output")
	}
	if !strings.Contains(string(content), `class="main-content"`) {
		t.Error("expected section class preserved")
	}
}

func TestBuildRealREADME(t *testing.T) {
	mdPath := "../../README.md"
	tmplPath := "../../config/website.tmpl.html"
	if _, err := os.Stat(mdPath); err != nil {
		t.Skip("README.md not found")
	}
	if _, err := os.Stat(tmplPath); err != nil {
		t.Skip("website template not found")
	}

	dir := t.TempDir()
	outPath := filepath.Join(dir, "index.html")

	if err := Build(mdPath, tmplPath, outPath); err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	info, err := os.Stat(outPath)
	if err != nil {
		t.Fatal(err)
	}
	if info.Size() < 10000 {
		t.Errorf("output too small: %d bytes", info.Size())
	}
	t.Logf("Generated %d bytes", info.Size())
}

func TestBuildFailsWithoutPlaceholder(t *testing.T) {
	dir := t.TempDir()

	mdPath := filepath.Join(dir, "README.md")
	if err := os.WriteFile(mdPath, []byte("# Title\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	tmplPath := filepath.Join(dir, "template.html")
	if err := os.WriteFile(tmplPath, []byte("<html><body><main></main></body></html>"), 0o644); err != nil {
		t.Fatal(err)
	}

	outPath := filepath.Join(dir, "index.html")
	err := Build(mdPath, tmplPath, outPath)
	if err == nil {
		t.Fatal("expected Build to fail when template has no supported placeholder")
	}
}

func TestBuildAddsHeadingIDs(t *testing.T) {
	dir := t.TempDir()

	md := "# Getting Started\n\n## Next Step\n"
	mdPath := filepath.Join(dir, "README.md")
	if err := os.WriteFile(mdPath, []byte(md), 0o644); err != nil {
		t.Fatal(err)
	}

	tmpl := `<!DOCTYPE html>
<html>
<body>
<section id="md" class="main-content"></section>
</body>
</html>`
	tmplPath := filepath.Join(dir, "template.html")
	if err := os.WriteFile(tmplPath, []byte(tmpl), 0o644); err != nil {
		t.Fatal(err)
	}

	outPath := filepath.Join(dir, "index.html")
	if err := Build(mdPath, tmplPath, outPath); err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	content, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatal(err)
	}

	html := string(content)
	if !strings.Contains(html, `id="getting-started"`) {
		t.Error("expected auto-generated heading id for h1")
	}
	if !strings.Contains(html, `id="next-step"`) {
		t.Error("expected auto-generated heading id for h2")
	}
}
