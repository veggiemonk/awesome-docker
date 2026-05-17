# Agent Guidelines for awesome-docker

This is a curated list of projects for Docker.
**The projects has to be for Docker, not just using Docker.**

## Commands

- Build CLI: `make build` (or `go build -o awesome-docker ./cmd/awesome-docker`)
- Rebuild from scratch: `make rebuild`
- Show local workflows: `make help`
- Format Go code: `make fmt`
- Run tests: `make test` (runs `go test ./internal/... -v`)
- Race tests: `make test-race`
- Lint README rules: `make lint` (runs `./awesome-docker lint`)
- Auto-fix lint issues: `make lint-fix`
- Check links: `make check` (runs `./awesome-docker check`; `GITHUB_TOKEN` enables GitHub repo checks)
- PR-safe link checks: `make check-pr`
- PR validation: `make validate` (lint + external link checks in PR mode)
- Build website: `make website` (generates `website/index.html` from `README.md`)
- Health scoring: `make health` (requires `GITHUB_TOKEN`, refreshes `config/health_cache.yaml`)
- Print health report (Markdown): `make report`
- Print health report (JSON): `make report-json` or `./awesome-docker report --json`
- Generate report files: `make report-file` (`HEALTH_REPORT.md`) and `make report-json-file` (`HEALTH_REPORT.json`)
- Maintenance shortcut: `make workflow-maint` (health + JSON report file)

## Architecture

- **Main content**: `README.md` (curated Docker/container resources)
- **CLI entrypoint**: `cmd/awesome-docker/main.go` (Cobra commands)
- **Core packages**:
    - `internal/parser` - parse README sections and entries
    - `internal/linter` - alphabetical/order/format validation + autofix
    - `internal/checker` - HTTP and GitHub link checks
    - `internal/scorer` - repository health scoring and report generation
    - `internal/cache` - exclude list and health cache read/write
    - `internal/builder` - render README to website HTML from template
- **Config**:
    - `config/exclude.yaml` - known link-check exclusions
    - `config/website.tmpl.html` - HTML template for site generation
    - `config/health_cache.yaml` - persisted health scoring cache
- **Generated outputs**:
    - `awesome-docker` - compiled CLI binary
    - `website/index.html` - generated website
    - `HEALTH_REPORT.md` - generated markdown report
    - `HEALTH_REPORT.json` - generated JSON report

## Content Guidelines (from CONTRIBUTING.md)

- **The projects has to be for Docker, not just using Docker.**
- Use one link per entry
- Prefer project/repository URLs over marketing pages
- Keep entries alphabetically ordered within each section
- Keep descriptions concise and concrete
- Use `:yen:` only for paid/commercial services
- Use `:ice_cube:` for stale projects (2+ years inactive)
- Remove archived/deprecated projects instead of tagging them
- Avoid duplicate links and redirect variants
