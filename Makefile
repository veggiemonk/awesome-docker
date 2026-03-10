SHELL := /bin/bash

BINARY ?= awesome-docker
GO ?= go
CMD_PACKAGE := ./cmd/awesome-docker
INTERNAL_PACKAGES := ./internal/...
WEBSITE_OUTPUT := website/index.html
HEALTH_CACHE := config/health_cache.yaml
HEALTH_REPORT_MD := HEALTH_REPORT.md
HEALTH_REPORT_JSON := HEALTH_REPORT.json

GO_SOURCES := $(shell find cmd internal -type f -name '*.go')
BUILD_INPUTS := $(GO_SOURCES) go.mod go.sum
WEBSITE_INPUTS := README.md config/website.tmpl.html
HEALTH_INPUTS := README.md config/exclude.yaml

.DEFAULT_GOAL := help

.PHONY: help \
	build rebuild clean \
	fmt test test-race \
	lint lint-fix check check-pr validate website \
	guard-github-token health health-cache \
	report report-json report-file report-json-file health-report \
	workflow-dev workflow-pr workflow-maint workflow-ci

help: ## Show the full local workflow and available targets
	@echo "awesome-docker Makefile"
	@echo
	@echo "Workflows:"
	@echo "  make workflow-dev    # local iteration (fmt + test + lint + check-pr + website)"
	@echo "  make workflow-pr     # recommended before opening/updating a PR"
	@echo "  make workflow-maint  # repository maintenance (health + JSON report)"
	@echo "  make workflow-ci     # CI-equivalent checks"
	@echo
	@echo "Core targets:"
	@echo "  make build           # build CLI binary"
	@echo "  make test            # run internal Go tests"
	@echo "  make lint            # validate README formatting/content rules"
	@echo "  make check           # check links (uses GITHUB_TOKEN when set)"
	@echo "  make validate        # run PR validation (lint + check --pr)"
	@echo "  make website         # generate website/index.html"
	@echo "  make report-file     # generate HEALTH_REPORT.md"
	@echo "  make report-json-file# generate HEALTH_REPORT.json"
	@echo "  make health          # refresh health cache (requires GITHUB_TOKEN)"
	@echo "  make report          # print markdown health report"
	@echo "  make report-json     # print full JSON health report"
	@echo
	@echo "Generated artifacts:"
	@echo "  $(BINARY)"
	@echo "  $(WEBSITE_OUTPUT)"
	@echo "  $(HEALTH_CACHE)"
	@echo "  $(HEALTH_REPORT_MD)"
	@echo "  $(HEALTH_REPORT_JSON)"

$(BINARY): $(BUILD_INPUTS)
	$(GO) build -o $(BINARY) $(CMD_PACKAGE)

build: $(BINARY) ## Build CLI binary

rebuild: clean build ## Rebuild from scratch

clean: ## Remove generated binary
	rm -f $(BINARY) $(HEALTH_REPORT_MD) $(HEALTH_REPORT_JSON)

fmt: ## Format Go code
	$(GO) fmt ./...

test: ## Run internal unit tests
	$(GO) test $(INTERNAL_PACKAGES) -v

test-race: ## Run internal tests with race detector
	$(GO) test $(INTERNAL_PACKAGES) -race

lint: build ## Validate README formatting/content rules
	./$(BINARY) lint

lint-fix: build ## Auto-fix lint issues when possible
	./$(BINARY) lint --fix

check: build ## Check links (GitHub checks enabled when GITHUB_TOKEN is set)
	./$(BINARY) check

check-pr: build ## Check links in PR mode (external links only)
	./$(BINARY) check --pr

validate: build ## Run PR validation (lint + check --pr)
	./$(BINARY) validate

$(WEBSITE_OUTPUT): $(BINARY) $(WEBSITE_INPUTS)
	./$(BINARY) build

website: $(WEBSITE_OUTPUT) ## Generate website from README

guard-github-token:
	@if [ -z "$$GITHUB_TOKEN" ]; then \
		echo "GITHUB_TOKEN is required for this target."; \
		echo "Set it with: export GITHUB_TOKEN=<token>"; \
		exit 1; \
	fi

$(HEALTH_CACHE): guard-github-token $(BINARY) $(HEALTH_INPUTS)
	./$(BINARY) health

health-cache: $(HEALTH_CACHE) ## Update config/health_cache.yaml

health: ## Refresh health cache from GitHub metadata
	@$(MAKE) --no-print-directory -B health-cache

report: build ## Print markdown health report from cache
	./$(BINARY) report

report-json: build ## Print full health report as JSON
	./$(BINARY) report --json

$(HEALTH_REPORT_MD): $(BINARY) $(HEALTH_CACHE)
	./$(BINARY) report > $(HEALTH_REPORT_MD)

report-file: $(HEALTH_REPORT_MD) ## Generate HEALTH_REPORT.md from cache

$(HEALTH_REPORT_JSON): $(BINARY) $(HEALTH_CACHE)
	./$(BINARY) report --json > $(HEALTH_REPORT_JSON)

report-json-file: $(HEALTH_REPORT_JSON) ## Generate HEALTH_REPORT.json from cache

health-report: health report-file ## Refresh health cache then generate HEALTH_REPORT.md

browse: build ## Launch interactive TUI browser
	./$(BINARY) browse

workflow-dev: fmt test lint check-pr website ## Full local development workflow

workflow-pr: fmt test validate ## Recommended workflow before opening a PR

workflow-maint: health report-json-file ## Weekly maintenance workflow

workflow-ci: test validate ## CI-equivalent validation workflow

update-ga:
	ratchet upgrade .github/workflows/*

update-go:
	go get -u go@latest
	go get -u ./...
	go mod tidy 
