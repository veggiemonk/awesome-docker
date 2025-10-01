# 🔧 Maintenance Guide for Awesome Docker

This guide helps maintainers keep the awesome-docker list up-to-date and high-quality.

## 🤖 Automated Systems

### Weekly Health Reports
- **What**: Checks all GitHub repositories for activity, archived status, and maintenance
- **When**: Every Monday at 9 AM UTC
- **Where**: Creates/updates a GitHub issue with label `health-report`
- **Action**: Review the report and mark abandoned projects with `:skull:`

### Broken Links Detection  
- **What**: Tests all links in README.md for availability
- **When**: Every Saturday at 2 AM UTC + on every PR
- **Where**: Creates/updates a GitHub issue with label `broken-links`
- **Action**: Fix or remove broken links, or add to exclusion list

### PR Validation
- **What**: Checks for duplicate links and basic validation
- **When**: On every pull request
- **Action**: Automated - contributors see results immediately

## 📋 Manual Maintenance Tasks

### Monthly Review (First Monday of the month)
1. Check health report issue for archived/stale projects
2. Mark archived projects with `:skull:` in README.md
3. Review projects with 2+ years of inactivity
4. Remove projects that are truly abandoned/broken

### Quarterly Deep Dive (Every 3 months)
1. Run: `npm run health-check` for detailed report
2. Review project categories - are they still relevant?
3. Check for popular new Docker tools to add
4. Update documentation links if newer versions exist

### Annual Cleanup (January)
1. Remove all `:skull:` projects older than 1 year
2. Review CONTRIBUTING.md guidelines
3. Update year references in documentation
4. Check Node.js version requirements

## 🛠️ Maintenance Commands

```bash
# Test all links (requires GITHUB_TOKEN)
npm test

# Test PR changes only
npm run test-pr

# Generate health report (requires GITHUB_TOKEN)
npm run health-check

# Build the website
npm run build

# Update dependencies
npm update
```

## 📊 Quality Standards

### Adding New Projects
- Must have clear documentation (README with install/usage)
- Should have activity within last 18 months
- GitHub project preferred over website links
- Must be Docker/container-related

### Marking Projects as Abandoned
Use `:skull:` emoji when:
- Repository is archived on GitHub
- No commits for 2+ years
- Project explicitly states it's deprecated
- Maintainer confirms abandonment

### Removing Projects
Only remove (don't just mark `:skull:`):
- Broken/404 links that can't be fixed
- Duplicate entries
- Spam or malicious projects
- Projects that never met quality standards

## 🚨 Emergency Procedures

### Critical Broken Links
If important resources are down:
1. Check if they moved (update URL)
2. Search for alternatives
3. Check Internet Archive for mirrors
4. Temporarily comment out until resolved

### Spam Pull Requests
1. Close immediately
2. Mark as spam
3. Block user if repeated offense
4. Don't engage in comments

## 📈 Metrics to Track

- Total projects: ~731 GitHub repos
- Health status: aim for <5% archived
- Link availability: aim for >98% working
- PR merge time: aim for <7 days
- Weekly contributor engagement

## 🤝 Getting Help

- Open a discussion in GitHub Discussions
- Check AGENTS.md for AI assistant guidelines
- Review CONTRIBUTING.md for contributor info

---

*Last updated: 2025-10-01*
