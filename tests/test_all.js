const fs = require('fs-extra');
const fetch = require('node-fetch');
const helper = require('./common');

function envvar_undefined(variable_name) {
    throw new Error(`${variable_name} must be defined`);
}

console.log({
    DEBUG: process.env.DEBUG || false,
});

const README = 'README.md';
const GITHUB_GQL_API = 'https://api.github.com/graphql';
const TOKEN = process.env.GITHUB_TOKEN || envvar_undefined('GITHUB_TOKEN');

const Authorization = `token ${TOKEN}`;

const make_GQL_options = (query) => ({
    method: 'POST',
    headers: {
        Authorization,
        'Content-Type': 'application/json',
        'user-agent':
            'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36',
    },
    body: JSON.stringify({ query }),
});

const extract_repos = (arr) =>
    arr
        .map((e) => e.substr('https://github.com/'.length).split('/'))
        .filter((r) => r.length === 2 && r[1] !== '');

const generate_GQL_query = (arr) =>
    `query AWESOME_REPOS{ ${arr
        .map(
            ([owner, name]) =>
                `repo_${owner.replace(/(-|\.)/g, '_')}_${name.replace(
                    /(-|\.)/g,
                    '_',
                )}: repository(owner: "${owner}", name:"${name}"){ nameWithOwner } `,
        )
        .join('')} }`;

async function main() {
    const has_error = {
        show: false,
        duplicates: '',
        other_links_error: '',
        github_repos: '',
    };
    const markdown = await fs.readFile(README, 'utf8');
    let links = helper.extract_all_links(markdown);
    links = links.filter((l) => !helper.exclude_from_list(l)); // exclude websites
    helper.LOG.debug_string({ links });

    console.log(`total links to check ${links.length}`);

    console.log('checking for duplicates links...');

    const duplicates = helper.find_duplicates(links);
    if (duplicates.length > 0) {
        has_error.show = true;
        has_error.duplicates = duplicates;
    }
    helper.LOG.debug_string({ duplicates });
    const [github_links, external_links] = helper.partition(links, (link) =>
        link.startsWith('https://github.com'),
    );

    console.log(`checking ${external_links.length} external links...`);

    const external_links_error = await helper.batch_fetch({
        arr: external_links,
        get: helper.fetch_link,
        post_filter_func: (x) => !x[1].ok,
        BATCH_SIZE: 8,
    });
    if (external_links_error.length > 0) {
        has_error.show = true;
        has_error.other_links_error = external_links_error;
    }

    console.log(`checking ${github_links.length} GitHub repositories...`);

    const repos = extract_repos(github_links);
    const query = generate_GQL_query(repos);
    const options = make_GQL_options(query);
    const gql_response = await fetch(GITHUB_GQL_API, options).then((r) =>
        r.json(),
    );
    if (gql_response.errors) {
        has_error.show = true;
        has_error.github_repos = gql_response.errors;
    }

    console.log({
        TEST_PASSED: has_error.show,
        GITHUB_REPOSITORY: github_links.length,
        EXTERNAL_LINKS: external_links.length,
    });

    if (has_error.show) {
        helper.LOG.error_string(has_error);
        process.exit(1);
    }
}

console.log('starting...');
main();
