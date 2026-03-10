package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/veggiemonk/awesome-docker/internal/builder"
	"github.com/veggiemonk/awesome-docker/internal/cache"
	"github.com/veggiemonk/awesome-docker/internal/checker"
	"github.com/veggiemonk/awesome-docker/internal/linter"
	"github.com/veggiemonk/awesome-docker/internal/parser"
	"github.com/veggiemonk/awesome-docker/internal/scorer"
	"github.com/veggiemonk/awesome-docker/internal/tui"
)

const (
	readmePath      = "README.md"
	excludePath     = "config/exclude.yaml"
	templatePath    = "config/website.tmpl.html"
	healthCachePath = "config/health_cache.yaml"
	websiteOutput   = "website/index.html"
	version         = "0.1.0"
)

type checkSummary struct {
	ExternalTotal int
	GitHubTotal   int
	Broken        []checker.LinkResult
	Redirected    []checker.LinkResult
	GitHubErrors  []error
	GitHubSkipped bool
}

func main() {
	root := &cobra.Command{
		Use:   "awesome-docker",
		Short: "Quality tooling for the awesome-docker curated list",
	}

	root.AddCommand(
		versionCmd(),
		lintCmd(),
		checkCmd(),
		healthCmd(),
		buildCmd(),
		reportCmd(),
		validateCmd(),
		ciCmd(),
		browseCmd(),
	)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run:   func(cmd *cobra.Command, args []string) { fmt.Printf("awesome-docker v%s\n", version) },
	}
}

func parseReadme() (parser.Document, error) {
	f, err := os.Open(readmePath)
	if err != nil {
		return parser.Document{}, err
	}
	defer f.Close()
	return parser.Parse(f)
}

func collectURLs(sections []parser.Section, urls *[]string) {
	for _, s := range sections {
		for _, e := range s.Entries {
			*urls = append(*urls, e.URL)
		}
		collectURLs(s.Children, urls)
	}
}

type entryMeta struct {
	Category    string
	Description string
}

func collectEntriesWithCategory(sections []parser.Section, parentPath string, out map[string]entryMeta) {
	for _, s := range sections {
		path := s.Title
		if parentPath != "" {
			path = parentPath + " > " + s.Title
		}
		for _, e := range s.Entries {
			out[e.URL] = entryMeta{Category: path, Description: e.Description}
		}
		collectEntriesWithCategory(s.Children, path, out)
	}
}

func runLinkChecks(prMode bool) (checkSummary, error) {
	doc, err := parseReadme()
	if err != nil {
		return checkSummary{}, fmt.Errorf("parse: %w", err)
	}

	var urls []string
	collectURLs(doc.Sections, &urls)

	exclude, err := cache.LoadExcludeList(excludePath)
	if err != nil {
		return checkSummary{}, fmt.Errorf("load exclude list: %w", err)
	}

	ghURLs, extURLs := checker.PartitionLinks(urls)

	summary := checkSummary{
		ExternalTotal: len(extURLs),
		GitHubTotal:   len(ghURLs),
	}

	results := checker.CheckLinks(extURLs, 10, exclude)
	for _, r := range results {
		if !r.OK {
			summary.Broken = append(summary.Broken, r)
		}
		if r.Redirected {
			summary.Redirected = append(summary.Redirected, r)
		}
	}

	if prMode {
		summary.GitHubSkipped = true
		return summary, nil
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		summary.GitHubSkipped = true
		return summary, nil
	}

	gc := checker.NewGitHubChecker(token)
	_, errs := gc.CheckRepos(context.Background(), ghURLs, 50)
	summary.GitHubErrors = errs
	return summary, nil
}

func runHealth(ctx context.Context) error {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return fmt.Errorf("GITHUB_TOKEN environment variable is required")
	}

	doc, err := parseReadme()
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	var urls []string
	collectURLs(doc.Sections, &urls)
	ghURLs, _ := checker.PartitionLinks(urls)

	fmt.Printf("Scoring %d GitHub repositories...\n", len(ghURLs))
	gc := checker.NewGitHubChecker(token)
	infos, errs := gc.CheckRepos(ctx, ghURLs, 50)
	for _, e := range errs {
		fmt.Printf("  error: %v\n", e)
	}
	if len(infos) == 0 {
		if len(errs) > 0 {
			return fmt.Errorf("failed to fetch GitHub metadata for all repositories (%d errors); check network/DNS and GITHUB_TOKEN", len(errs))
		}
		return fmt.Errorf("no GitHub repositories found in README")
	}

	scored := scorer.ScoreAll(infos)

	meta := make(map[string]entryMeta)
	collectEntriesWithCategory(doc.Sections, "", meta)
	for i := range scored {
		if m, ok := meta[scored[i].URL]; ok {
			scored[i].Category = m.Category
			scored[i].Description = m.Description
		}
	}

	cacheEntries := scorer.ToCacheEntries(scored)

	hc, err := cache.LoadHealthCache(healthCachePath)
	if err != nil {
		return fmt.Errorf("load cache: %w", err)
	}
	hc.Merge(cacheEntries)
	if err := cache.SaveHealthCache(healthCachePath, hc); err != nil {
		return fmt.Errorf("save cache: %w", err)
	}

	fmt.Printf("Cache updated: %d entries in %s\n", len(hc.Entries), healthCachePath)
	return nil
}

func scoredFromCache() ([]scorer.ScoredEntry, error) {
	hc, err := cache.LoadHealthCache(healthCachePath)
	if err != nil {
		return nil, fmt.Errorf("load cache: %w", err)
	}
	if len(hc.Entries) == 0 {
		return nil, fmt.Errorf("no cache data, run 'health' first")
	}

	scored := make([]scorer.ScoredEntry, 0, len(hc.Entries))
	for _, e := range hc.Entries {
		scored = append(scored, scorer.ScoredEntry{
			URL:         e.URL,
			Name:        e.Name,
			Status:      scorer.Status(e.Status),
			Stars:       e.Stars,
			Forks:       e.Forks,
			HasLicense:  e.HasLicense,
			LastPush:    e.LastPush,
			Category:    e.Category,
			Description: e.Description,
		})
	}
	return scored, nil
}

func markdownReportFromCache() (string, error) {
	scored, err := scoredFromCache()
	if err != nil {
		return "", err
	}
	return scorer.GenerateReport(scored), nil
}

func writeGitHubOutput(path, key, value string) error {
	if path == "" {
		return nil
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open github output file: %w", err)
	}
	defer f.Close()
	if _, err := fmt.Fprintf(f, "%s=%s\n", key, value); err != nil {
		return fmt.Errorf("write github output: %w", err)
	}
	return nil
}

func sanitizeOutputValue(v string) string {
	v = strings.ReplaceAll(v, "\n", " ")
	v = strings.ReplaceAll(v, "\r", " ")
	return strings.TrimSpace(v)
}

func buildBrokenLinksIssueBody(summary checkSummary, runErr error) string {
	var b strings.Builder
	b.WriteString("# Broken Links Detected\n\n")

	if runErr != nil {
		b.WriteString("The link checker failed to execute cleanly.\n\n")
		b.WriteString("## Failure\n\n")
		fmt.Fprintf(&b, "- %s\n\n", runErr)
	} else {
		fmt.Fprintf(&b, "- Broken links: %d\n", len(summary.Broken))
		fmt.Fprintf(&b, "- Redirected links: %d\n", len(summary.Redirected))
		fmt.Fprintf(&b, "- GitHub API errors: %d\n\n", len(summary.GitHubErrors))

		if len(summary.Broken) > 0 {
			b.WriteString("## Broken Links\n\n")
			for _, r := range summary.Broken {
				fmt.Fprintf(&b, "- `%s` -> `%d %s`\n", r.URL, r.StatusCode, strings.TrimSpace(r.Error))
			}
			b.WriteString("\n")
		}

		if len(summary.GitHubErrors) > 0 {
			b.WriteString("## GitHub API Errors\n\n")
			for _, e := range summary.GitHubErrors {
				fmt.Fprintf(&b, "- `%s`\n", e)
			}
			b.WriteString("\n")
		}
	}

	b.WriteString("## Action Required\n\n")
	b.WriteString("- Update the URL if the resource moved\n")
	b.WriteString("- Remove the entry if permanently unavailable\n")
	b.WriteString("- Add to `config/exclude.yaml` if a known false positive\n")
	b.WriteString("- Investigate GitHub API/auth failures when present\n\n")
	b.WriteString("---\n")
	b.WriteString("*Auto-generated by awesome-docker ci broken-links*\n")
	return b.String()
}

func buildHealthReportIssueBody(report string, healthErr error) string {
	var b strings.Builder
	if healthErr != nil {
		b.WriteString("WARNING: health refresh failed in this run; showing latest cached report.\n\n")
		fmt.Fprintf(&b, "Error: `%s`\n\n", healthErr)
	}
	b.WriteString(report)
	if !strings.HasSuffix(report, "\n") {
		b.WriteString("\n")
	}
	b.WriteString("\n---\n")
	b.WriteString("*Auto-generated weekly by awesome-docker ci health-report*\n")
	return b.String()
}

func lintCmd() *cobra.Command {
	var fix bool
	cmd := &cobra.Command{
		Use:   "lint",
		Short: "Validate README formatting",
		RunE: func(cmd *cobra.Command, args []string) error {
			doc, err := parseReadme()
			if err != nil {
				return fmt.Errorf("parse: %w", err)
			}

			result := linter.Lint(doc)
			for _, issue := range result.Issues {
				fmt.Println(issue)
			}

			if result.Errors > 0 {
				fmt.Printf("\n%d errors, %d warnings\n", result.Errors, result.Warnings)
				if !fix {
					return fmt.Errorf("lint failed with %d errors", result.Errors)
				}
				count, err := linter.FixFile(readmePath)
				if err != nil {
					return fmt.Errorf("fix: %w", err)
				}
				fmt.Printf("Fixed %d lines in %s\n", count, readmePath)
			} else {
				fmt.Printf("OK: %d warnings\n", result.Warnings)
			}

			return nil
		},
	}
	cmd.Flags().BoolVar(&fix, "fix", false, "Auto-fix formatting issues")
	return cmd
}

func checkCmd() *cobra.Command {
	var prMode bool
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check links for reachability",
		RunE: func(cmd *cobra.Command, args []string) error {
			summary, err := runLinkChecks(prMode)
			if err != nil {
				return err
			}

			fmt.Printf("Checking %d external links...\n", summary.ExternalTotal)
			if !prMode {
				if summary.GitHubSkipped {
					fmt.Println("GITHUB_TOKEN not set, skipping GitHub repo checks")
				} else {
					fmt.Printf("Checking %d GitHub repositories...\n", summary.GitHubTotal)
				}
			}

			for _, e := range summary.GitHubErrors {
				fmt.Printf("  GitHub error: %v\n", e)
			}

			if len(summary.Redirected) > 0 {
				fmt.Printf("\n%d redirected links (consider updating):\n", len(summary.Redirected))
				for _, r := range summary.Redirected {
					fmt.Printf("  %s -> %s\n", r.URL, r.RedirectURL)
				}
			}

			if len(summary.Broken) > 0 {
				fmt.Printf("\n%d broken links:\n", len(summary.Broken))
				for _, r := range summary.Broken {
					fmt.Printf("  %s -> %d %s\n", r.URL, r.StatusCode, r.Error)
				}
			}
			if len(summary.Broken) > 0 && len(summary.GitHubErrors) > 0 {
				return fmt.Errorf("found %d broken links and %d GitHub API errors", len(summary.Broken), len(summary.GitHubErrors))
			}
			if len(summary.Broken) > 0 {
				return fmt.Errorf("found %d broken links", len(summary.Broken))
			}
			if len(summary.GitHubErrors) > 0 {
				return fmt.Errorf("github checks failed with %d errors", len(summary.GitHubErrors))
			}

			fmt.Println("All links OK")
			return nil
		},
	}
	cmd.Flags().BoolVar(&prMode, "pr", false, "PR mode: skip GitHub API checks")
	return cmd
}

func healthCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "health",
		Short: "Score repository health and update cache",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runHealth(context.Background())
		},
	}
}

func buildCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "build",
		Short: "Generate website from README",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := builder.Build(readmePath, templatePath, websiteOutput); err != nil {
				return err
			}
			fmt.Printf("Website built: %s\n", websiteOutput)
			return nil
		},
	}
}

func reportCmd() *cobra.Command {
	var jsonOutput bool
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Generate health report from cache",
		RunE: func(cmd *cobra.Command, args []string) error {
			scored, err := scoredFromCache()
			if err != nil {
				return err
			}

			if jsonOutput {
				payload, err := scorer.GenerateJSONReport(scored)
				if err != nil {
					return fmt.Errorf("json report: %w", err)
				}
				fmt.Println(string(payload))
				return nil
			}

			report := scorer.GenerateReport(scored)
			fmt.Print(report)
			return nil
		},
	}

	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Output full health report as JSON")
	return cmd
}

func validateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "validate",
		Short: "PR validation: lint + check --pr",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("=== Linting ===")
			doc, err := parseReadme()
			if err != nil {
				return fmt.Errorf("parse: %w", err)
			}

			result := linter.Lint(doc)
			for _, issue := range result.Issues {
				fmt.Println(issue)
			}
			if result.Errors > 0 {
				fmt.Printf("\n%d errors, %d warnings\n", result.Errors, result.Warnings)
				return fmt.Errorf("lint failed with %d errors", result.Errors)
			}
			fmt.Printf("Lint OK: %d warnings\n", result.Warnings)

			fmt.Println("\n=== Checking links (PR mode) ===")
			summary, err := runLinkChecks(true)
			if err != nil {
				return err
			}
			fmt.Printf("Checking %d external links...\n", summary.ExternalTotal)
			if len(summary.Broken) > 0 {
				fmt.Printf("\n%d broken links:\n", len(summary.Broken))
				for _, r := range summary.Broken {
					fmt.Printf("  %s -> %d %s\n", r.URL, r.StatusCode, r.Error)
				}
				return fmt.Errorf("found %d broken links", len(summary.Broken))
			}

			fmt.Println("\nValidation passed")
			return nil
		},
	}
}

func ciCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ci",
		Short: "CI-oriented helper commands",
	}
	cmd.AddCommand(
		ciBrokenLinksCmd(),
		ciHealthReportCmd(),
	)
	return cmd
}

func ciBrokenLinksCmd() *cobra.Command {
	var issueFile string
	var githubOutput string
	var strict bool

	cmd := &cobra.Command{
		Use:   "broken-links",
		Short: "Run link checks and emit CI outputs/artifacts",
		RunE: func(cmd *cobra.Command, args []string) error {
			summary, runErr := runLinkChecks(false)

			hasErrors := runErr != nil || len(summary.Broken) > 0 || len(summary.GitHubErrors) > 0
			exitCode := 0
			if hasErrors {
				exitCode = 1
			}
			if runErr != nil {
				exitCode = 2
			}

			if issueFile != "" && hasErrors {
				body := buildBrokenLinksIssueBody(summary, runErr)
				if err := os.WriteFile(issueFile, []byte(body), 0o644); err != nil {
					return fmt.Errorf("write issue file: %w", err)
				}
			}

			if err := writeGitHubOutput(githubOutput, "has_errors", strconv.FormatBool(hasErrors)); err != nil {
				return err
			}
			if err := writeGitHubOutput(githubOutput, "check_exit_code", strconv.Itoa(exitCode)); err != nil {
				return err
			}
			if err := writeGitHubOutput(githubOutput, "broken_count", strconv.Itoa(len(summary.Broken))); err != nil {
				return err
			}
			if err := writeGitHubOutput(githubOutput, "github_error_count", strconv.Itoa(len(summary.GitHubErrors))); err != nil {
				return err
			}
			if runErr != nil {
				if err := writeGitHubOutput(githubOutput, "run_error", sanitizeOutputValue(runErr.Error())); err != nil {
					return err
				}
			}

			if runErr != nil {
				fmt.Printf("CI broken-links run error: %v\n", runErr)
			}
			if hasErrors {
				fmt.Printf("CI broken-links found %d broken links and %d GitHub errors\n", len(summary.Broken), len(summary.GitHubErrors))
			} else {
				fmt.Println("CI broken-links found no errors")
			}

			if strict {
				if runErr != nil {
					return runErr
				}
				if hasErrors {
					return fmt.Errorf("found %d broken links and %d GitHub API errors", len(summary.Broken), len(summary.GitHubErrors))
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&issueFile, "issue-file", "broken_links_issue.md", "Path to write issue markdown body")
	cmd.Flags().StringVar(&githubOutput, "github-output", "", "Path to GitHub output file (typically $GITHUB_OUTPUT)")
	cmd.Flags().BoolVar(&strict, "strict", false, "Return non-zero when errors are found")
	return cmd
}

func ciHealthReportCmd() *cobra.Command {
	var issueFile string
	var githubOutput string
	var strict bool

	cmd := &cobra.Command{
		Use:   "health-report",
		Short: "Refresh health cache, render report, and emit CI outputs/artifacts",
		RunE: func(cmd *cobra.Command, args []string) error {
			healthErr := runHealth(context.Background())
			report, reportErr := markdownReportFromCache()

			healthOK := healthErr == nil
			reportOK := reportErr == nil
			hasReport := reportOK && strings.TrimSpace(report) != ""
			hasErrors := !healthOK || !reportOK

			if hasReport && issueFile != "" {
				body := buildHealthReportIssueBody(report, healthErr)
				if err := os.WriteFile(issueFile, []byte(body), 0o644); err != nil {
					return fmt.Errorf("write issue file: %w", err)
				}
			}

			if err := writeGitHubOutput(githubOutput, "has_report", strconv.FormatBool(hasReport)); err != nil {
				return err
			}
			if err := writeGitHubOutput(githubOutput, "health_ok", strconv.FormatBool(healthOK)); err != nil {
				return err
			}
			if err := writeGitHubOutput(githubOutput, "report_ok", strconv.FormatBool(reportOK)); err != nil {
				return err
			}
			if err := writeGitHubOutput(githubOutput, "has_errors", strconv.FormatBool(hasErrors)); err != nil {
				return err
			}
			if healthErr != nil {
				if err := writeGitHubOutput(githubOutput, "health_error", sanitizeOutputValue(healthErr.Error())); err != nil {
					return err
				}
			}
			if reportErr != nil {
				if err := writeGitHubOutput(githubOutput, "report_error", sanitizeOutputValue(reportErr.Error())); err != nil {
					return err
				}
			}

			if healthErr != nil {
				fmt.Printf("CI health-report health error: %v\n", healthErr)
			}
			if reportErr != nil {
				fmt.Printf("CI health-report report error: %v\n", reportErr)
			}
			if hasReport {
				fmt.Println("CI health-report generated report artifact")
			} else {
				fmt.Println("CI health-report has no report artifact")
			}

			if strict {
				if healthErr != nil {
					return healthErr
				}
				if reportErr != nil {
					return reportErr
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&issueFile, "issue-file", "health_report.txt", "Path to write health issue markdown body")
	cmd.Flags().StringVar(&githubOutput, "github-output", "", "Path to GitHub output file (typically $GITHUB_OUTPUT)")
	cmd.Flags().BoolVar(&strict, "strict", false, "Return non-zero when health/report fails")
	return cmd
}

func browseCmd() *cobra.Command {
	var cachePath string
	cmd := &cobra.Command{
		Use:   "browse",
		Short: "Interactive TUI browser for awesome-docker resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			hc, err := cache.LoadHealthCache(cachePath)
			if err != nil {
				return fmt.Errorf("load cache: %w", err)
			}
			if len(hc.Entries) == 0 {
				return fmt.Errorf("no cache data; run 'awesome-docker health' first")
			}
			return tui.Run(hc.Entries)
		},
	}
	cmd.Flags().StringVar(&cachePath, "cache", healthCachePath, "Path to health cache YAML")
	return cmd
}
