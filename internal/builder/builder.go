package builder

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Build converts a Markdown file to HTML using a template.
// The template must contain a placeholder element that will be replaced with the rendered content.
func Build(markdownPath, templatePath, outputPath string) error {
	md, err := os.ReadFile(markdownPath)
	if err != nil {
		return fmt.Errorf("read markdown: %w", err)
	}

	tmpl, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("read template: %w", err)
	}

	// Convert markdown to HTML
	gm := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)
	var buf bytes.Buffer
	if err := gm.Convert(md, &buf); err != nil {
		return fmt.Errorf("convert markdown: %w", err)
	}

	// Inject into template — support both placeholder formats
	output := string(tmpl)
	replacements := []struct {
		old string
		new string
	}{
		{
			old: `<div id="md"></div>`,
			new: `<div id="md">` + buf.String() + `</div>`,
		},
		{
			old: `<section id="md" class="main-content"></section>`,
			new: `<section id="md" class="main-content">` + buf.String() + `</section>`,
		},
	}

	replaced := false
	for _, r := range replacements {
		if strings.Contains(output, r.old) {
			output = strings.Replace(output, r.old, r.new, 1)
			replaced = true
			break
		}
	}
	if !replaced {
		return fmt.Errorf("template missing supported markdown placeholder")
	}

	if err := os.WriteFile(outputPath, []byte(output), 0o644); err != nil {
		return fmt.Errorf("write output: %w", err)
	}
	return nil
}
