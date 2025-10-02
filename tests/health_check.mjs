import fs from 'fs-extra';
import fetch from 'node-fetch';
import helper from './common.mjs';

const README = 'README.md';
const GITHUB_GQL_API = 'https://api.github.com/graphql';
const TOKEN = process.env.GITHUB_TOKEN || '';

if (!TOKEN) {
    console.error('GITHUB_TOKEN environment variable is required');
    process.exit(1);
}

const Authorization = `token ${TOKEN}`;

const LOG = {
    info: (...args) => console.log('ℹ️ ', ...args),
    warn: (...args) => console.warn('⚠️ ', ...args),
    error: (...args) => console.error('❌', ...args),
};

// Extract GitHub repos from links
const extract_repos = (arr) =>
    arr
        .map((e) => e.substr('https://github.com/'.length).split('/'))
        .filter((r) => r.length === 2 && r[1] !== '');

// Generate GraphQL query to check repo health
const generate_health_query = (repos) => {
    const repoQueries = repos.map(([owner, name]) => {
        const safeName = `repo_${owner.replace(/(-|\.)/g, '_')}_${name.replace(/(-|\.)/g, '_')}`;
        return `${safeName}: repository(owner: "${owner}", name:"${name}"){
            nameWithOwner
            isArchived
            pushedAt
            createdAt
            stargazerCount
            forkCount
            isDisabled
            isFork
            isLocked
            isPrivate
        }`;
    }).join('\n');
    
    return `query REPO_HEALTH { ${repoQueries} }`;
};

// Batch repos into smaller chunks for GraphQL
function* batchRepos(repos, size = 50) {
    for (let i = 0; i < repos.length; i += size) {
        yield repos.slice(i, i + size);
    }
}

async function checkRepoHealth(repos) {
    const results = {
        archived: [],
        stale: [],      // No commits in 2+ years
        inactive: [],   // No commits in 1-2 years
        healthy: [],
        disabled: [],
        total: repos.length,
    };

    const twoYearsAgo = new Date();
    twoYearsAgo.setFullYear(twoYearsAgo.getFullYear() - 2);
    
    const oneYearAgo = new Date();
    oneYearAgo.setFullYear(oneYearAgo.getFullYear() - 1);

    LOG.info(`Checking health of ${repos.length} repositories...`);

    for (const batch of batchRepos(repos)) {
        const query = generate_health_query(batch);
        const options = {
            method: 'POST',
            headers: {
                Authorization,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ query }),
        };

        try {
            const response = await fetch(GITHUB_GQL_API, options);
            const data = await response.json();

            if (data.errors) {
                LOG.error('GraphQL errors:', data.errors);
                continue;
            }

            for (const [key, repo] of Object.entries(data.data)) {
                if (!repo) continue;

                const pushedAt = new Date(repo.pushedAt);
                const repoInfo = {
                    name: repo.nameWithOwner,
                    pushedAt: repo.pushedAt,
                    stars: repo.stargazerCount,
                    url: `https://github.com/${repo.nameWithOwner}`,
                };

                if (repo.isArchived) {
                    results.archived.push(repoInfo);
                } else if (repo.isDisabled) {
                    results.disabled.push(repoInfo);
                } else if (pushedAt < twoYearsAgo) {
                    results.stale.push(repoInfo);
                } else if (pushedAt < oneYearAgo) {
                    results.inactive.push(repoInfo);
                } else {
                    results.healthy.push(repoInfo);
                }
            }
        } catch (error) {
            LOG.error('Batch fetch error:', error.message);
        }

        // Rate limiting - wait a bit between batches
        await new Promise(resolve => setTimeout(resolve, 1000));
    }

    return results;
}

function generateReport(results) {
    const report = [];
    
    report.push('# 🏥 Awesome Docker - Health Check Report\n');
    report.push(`**Generated:** ${new Date().toISOString()}\n`);
    report.push(`**Total Repositories:** ${results.total}\n`);
    
    report.push('\n## 📊 Summary\n');
    report.push(`- ✅ Healthy (updated in last year): ${results.healthy.length}`);
    report.push(`- ⚠️  Inactive (1-2 years): ${results.inactive.length}`);
    report.push(`- 🪦 Stale (2+ years): ${results.stale.length}`);
    report.push(`- 📦 Archived: ${results.archived.length}`);
    report.push(`- 🚫 Disabled: ${results.disabled.length}\n`);

    if (results.archived.length > 0) {
        report.push('\n## 📦 Archived Repositories (Should mark as :skull:)\n');
        results.archived.forEach(repo => {
            report.push(`- [${repo.name}](${repo.url}) - ⭐ ${repo.stars} - Last push: ${repo.pushedAt}`);
        });
    }

    if (results.stale.length > 0) {
        report.push('\n## 🪦 Stale Repositories (No activity in 2+ years)\n');
        results.stale.slice(0, 50).forEach(repo => {
            report.push(`- [${repo.name}](${repo.url}) - ⭐ ${repo.stars} - Last push: ${repo.pushedAt}`);
        });
        if (results.stale.length > 50) {
            report.push(`\n... and ${results.stale.length - 50} more`);
        }
    }

    if (results.inactive.length > 0) {
        report.push('\n## ⚠️ Inactive Repositories (No activity in 1-2 years)\n');
        report.push('_These may still be stable/complete projects - review individually_\n');
        results.inactive.slice(0, 30).forEach(repo => {
            report.push(`- [${repo.name}](${repo.url}) - ⭐ ${repo.stars} - Last push: ${repo.pushedAt}`);
        });
        if (results.inactive.length > 30) {
            report.push(`\n... and ${results.inactive.length - 30} more`);
        }
    }

    return report.join('\n');
}

async function main() {
    const markdown = await fs.readFile(README, 'utf8');
    let links = helper.extract_all_links(markdown);
    
    const github_links = links.filter(link => 
        link.startsWith('https://github.com') && 
        !helper.exclude_from_list(link) &&
        !link.includes('/issues') &&
        !link.includes('/pull') &&
        !link.includes('/wiki') &&
        !link.includes('#')
    );

    const repos = extract_repos(github_links);
    const results = await checkRepoHealth(repos);
    
    const report = generateReport(results);
    
    // Save report
    await fs.writeFile('HEALTH_REPORT.md', report);
    LOG.info('Health report saved to HEALTH_REPORT.md');
    
    // Also print summary to console
    console.log('\n' + report);
    
    // Exit with error if there are actionable items
    if (results.archived.length > 0 || results.stale.length > 10) {
        LOG.warn(`Found ${results.archived.length} archived and ${results.stale.length} stale repos`);
        process.exit(1);
    }
}

console.log('Starting health check...');
main();
