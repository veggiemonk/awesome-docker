# Contributing to awesome-docker

Thanks for taking the time to contribute.

This repository is a curated list of Docker/container resources plus a Go-based maintenance CLI used by CI. Contributions are welcome for both content and tooling.

Please read and follow the [Code of Conduct](./CODE_OF_CONDUCT.md).

## What We Accept

- New high-quality Docker/container-related projects
- Fixes to descriptions, ordering, or categorization
- Removal of broken, archived, deprecated, or duplicate entries
- Improvements to the Go CLI and GitHub workflows

## README Entry Rules

- Use one link per entry.
- Prefer GitHub project/repository URLs over marketing pages.
- Keep entries alphabetically sorted within their section.
- Keep descriptions concise and concrete.
- Use `:yen:` for paid/commercial services.
- Use `:ice_cube:` for stale projects (2+ years inactive).
- Do not use `:skull:`; archived/deprecated projects should be removed.
- Avoid duplicate links and redirect variants.

## Local Validation

```bash
# Build CLI
make build

# Validate README formatting and content
make lint

# Run code tests (when touching Go code)
make test

# Optional: full external checks (requires GITHUB_TOKEN)
./awesome-docker check
./awesome-docker validate
```

## Pull Request Expectations

- Keep the PR focused to one logical change.
- Explain what changed and why.
- If adding entries, include the target category.
- If removing entries, explain why (archived, broken, duplicate, etc.).
- Fill in the PR template checklist.

## Maintainer Notes

- Changes should be reviewed before merge.
- Prefer helping contributors improve a PR over silently rejecting it.
- Keep `.github` documentation and workflows aligned with current tooling.
