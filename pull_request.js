const fs = require('fs-extra');
const fetch = require('node-fetch');
const exclude = require('./exclude_in_test.json');

const handleFailure = error => {
    console.error(`${error.message}: ${error.stack}`, { error });
    process.exit(1);
};

process.on('unhandledRejection', handleFailure);
const LOG = {
    error: (...args) => console.error('❌ ERROR', args),
    error_string: (...args) =>
        console.error('❌ ERROR', JSON.stringify({ ...args }, null, 2)),
    debug: (...args) => {
        if (process.env.DEBUG) console.log('>>> DEBUG: ', { ...args });
    },
    debug_string: (...args) => {
        if (process.env.DEBUG)
            console.log('>>> DEBUG: ', JSON.stringify({ ...args }, null, 2));
    },
};
function envvar_undefined(variable_name) {
    throw new Error(`${variable_name} must be defined`);
}

/** ------------------------------------------------------------------------------------
 * CONSTANTS
 */
const README = 'README.md';

const GITHUB_API = 'https://api.github.com';

const TOKEN = process.env.GITHUB_TOKEN || envvar_undefined('GITHUB_TOKEN');

const LINKS_OPTIONS = {
    // redirect: 'error',
    headers: {
        'Content-Type': 'application/json',
        'user-agent':
            'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36',
    },
};

const GITHUB_API_OPTIONS = {
    Authorization: `token ${TOKEN}`,
    headers: {
        Accept: 'application/vnd.github.v3+json',
        'user-agent':
            'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36',
    },
};

/** ------------------------------------------------------------------------------------
 *
 */

const extract_all_links = markdown => {
    // if you have a problem and you try to solve it with a regex,
    // now you have two problems
    // TODO: replace this mess with a mardown parser ?
    const re = /(((https:(?:\/\/)?)(?:[-;:&=+$,\w]+@)?[A-Za-z0-9.-]+|(?:www\.|[-;:&=+$,\w]+@)[A-Za-z0-9.-]+)((?:\/[+~%/.\w\-_]*)?\??(?:[-+=&;%@.\w_]*)#?(?:[.!/@\-\\\w]*))?)/g;
    return markdown.match(re);
};

const find_duplicates = arr => {
    const hm = {};
    const dup = [];
    arr.forEach(e => {
        if (hm[e]) dup.push(e);
        else hm[e] = true;
    });
    return dup;
};

const partition = (arr, func) => {
    const ap = [[], []];
    arr.forEach(e => (func(e) ? ap[0].push(e) : ap[1].push(e)));
    return ap;
};

const extract_repos = arr =>
    arr
        .map(e => e.substr('https://github.com/'.length).split('/'))
        .filter(r => r.length === 2 && r[1] !== '')
        .map(x => x.join('/'));

const exclude_length = exclude.length;
const exclude_from_list = link => {
    let is_excluded = false;
    for (let i = 0; i < exclude_length; i += 1) {
        if (link.startsWith(exclude[i])) {
            is_excluded = true;
            break;
        }
    }
    return is_excluded;
};

/** ------------------------------------------------------------------------------------
 *  FETCH
 */
async function fetch_link(url) {
    try {
        const { ok, statusText, redirected } = await fetch(url, LINKS_OPTIONS);
        return [url, { ok, status: statusText, redirected }];
    } catch (error) {
        return [url, { ok: false, status: error.message }];
    }
}

async function fetch_repo(repo) {
    try {
        const response = await fetch(
            `${GITHUB_API}/repos/${repo}`,
            GITHUB_API_OPTIONS,
        );
        const { ok, statusText, redirected, headers, url } = response;
        // const json = await response.json();
        if (redirected) {
            for (const pair of headers.entries()) {
                console.log(`${pair[0]}: ${pair[1]}`);
            }
            console.log(`redirected to ${url}`);
            const r = await fetch(url, GITHUB_API_OPTIONS);
            console.log({
                ok: r.ok,
                status: r.statusText,
                redirected: r.redirected,
                url: r.url,
            });
            for (const pair of r.headers.entries()) {
                console.log(`>>> REDIRECT: ${pair[0]}: ${pair[1]}`);
            }
        }
        return [repo, { ok, status: statusText, redirected, url }];
    } catch (error) {
        return [repo, { ok: false, status: error.message }];
    }
}

async function batch_fetch({ arr, get, post_filter_func, BATCH_SIZE = 8 }) {
    const result = [];
    /* eslint-disable no-await-in-loop */
    for (let i = 0; i < arr.length; i += BATCH_SIZE) {
        const batch = arr.slice(i, i + BATCH_SIZE);
        LOG.debug_string({ batch });
        let res = await Promise.all(batch.map(get));
        console.log(`batch fetched...${i + BATCH_SIZE}`);
        res = post_filter_func ? res.filter(post_filter_func) : res;
        LOG.debug_string({ res });
        result.push(...res);
    }
    return result;
}

/** ------------------------------------------------------------------------------------
 *  MAIN
 */
async function main() {
    const has_error = {
        show: false,
        duplicates: '',
        other_links_error: '',
        github_repos: '',
        // query: '',
    };
    const markdown = await fs.readFile(README, 'utf8');
    let links = extract_all_links(markdown);
    links = links.filter(l => !exclude_from_list(l)); // exclude websites
    LOG.debug_string({ links });

    console.log(`total links to check ${links.length}`);

    console.log('checking for duplicates links...');

    const duplicates = find_duplicates(links);
    if (duplicates.length > 0) {
        has_error.show = true;
        has_error.duplicates = duplicates;
    }
    LOG.debug_string({ duplicates });
    const [github_links, external_links] = partition(links, link =>
        link.startsWith('https://github.com'),
    );

    console.log(`checking ${links.length} links...`);

    const external_links_error = await batch_fetch({
        arr: external_links.slice(0, 10),
        get: fetch_link,
        post_filter_func: x => !x[1].ok,
        BATCH_SIZE: 32,
    });
    if (external_links_error.length > 0) {
        // LOG.debug({ external_links_error });
        has_error.show = true;
        has_error.other_links_error = external_links_error;
    }

    const github_links_error = await batch_fetch({
        arr: extract_repos(github_links).slice(0, 10),
        get: fetch_repo,
        // post_filter_func: x => !x[1].ok,
        BATCH_SIZE: 32,
    });
    if (github_links_error.length > 0) {
        // LOG.debug({ github_links_error });
        has_error.show = true;
        has_error.github_repos = github_links_error;
    }

    if (has_error.show) {
        LOG.error_string(has_error);
        process.exit(1);
    }
    console.log({
        TEST_PASSED: !has_error.show,
        GITHUB_REPOSITORY: github_links.length,
        EXTERNAL_LINKS: external_links.length,
        TOTAL_LINKS: links.length,
    });
}

console.log('starting...');
// eslint-disable-next-line no-unused-expressions
process.env.DEBUG && console.log('debugging mode on.');
main();
