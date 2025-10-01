# Agent Guidelines for awesome-docker

## Commands
- Build website: `npm run build` (converts README.md to website/index.html)
- Test all links: `npm test` (runs tests/test_all.mjs, requires GITHUB_TOKEN)
- Test PR changes: `npm run test-pr` (runs tests/pull_request.mjs, checks duplicates)
- Health check: `npm run health-check` (generates HEALTH_REPORT.md, requires GITHUB_TOKEN)

## Architecture
- **Main content**: README.md - curated list of Docker resources (markdown format)
- **Build script**: build.js - converts README.md to HTML using showdown & cheerio
- **Tests**: tests/*.mjs - link validation, duplicate detection, URL checking
- **Website**: website/ - static site deployment folder

## Code Style
- **Language**: Node.js with ES modules (.mjs) for tests, CommonJS for build.js
- **Imports**: Use ES6 imports in .mjs files, require() in .js files
- **Error handling**: Use try/catch with LOG.error() and process.exit(1) for failures
- **Logging**: Use LOG object with error/debug methods (see build.js for pattern)
- **Async**: Prefer async/await over callbacks

## Content Guidelines (from CONTRIBUTING.md)
- Link to GitHub projects, not websites
- Entries are listed alphabetically (from A to Z)
- Entries must be Docker/container-related with clear documentation
- Include project description, installation, and usage examples
- Mark WIP projects explicitly
- Avoid outdated tutorials/blog posts unless advanced/specific
