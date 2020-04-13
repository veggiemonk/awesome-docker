const fs = require('fs-extra');
const fetch = require('node-fetch');

function envvar_undefined(variable_name) {
  throw new Error(`${variable_name} must be defined`);
}
console.log({
  DEBUG: process.env.DEBUG,
});

const README = 'README.md';
const GITHUB_GQL_API = 'https://api.github.com/graphql';
const TOKEN = process.env.GITHUB_TOKEN || envvar_undefined('GITHUB_TOKEN');

const LINKS_OPTIONS = {
  redirect: 'error',
  headers: {
    'Content-Type': 'application/json',
    'user-agent':
      'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36',
  },
};

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

const LOG = {
  error: (...args) => console.error('❌ ERROR', args),
  error_string: (...args) =>
    console.error(
      '❌ ERROR',
      args.map((a) => JSON.stringify(a)),
    ),
  debug: (...args) => {
    if (process.env.DEBUG) console.log('>>> DEBUG: ', { ...args });
  },
  debug_string: (...args) => {
    if (process.env.DEBUG)
      console.log('>>> DEBUG: ', JSON.stringify({ ...args }));
  },
};

const handleFailure = (error) => {
  console.error(`${error.message}: ${error.stack}`, { error });
  process.exit(1);
};

process.on('unhandledRejection', handleFailure);

const extract_all_links = (markdown) => {
  // if you have a problem and you try to solve it with a regex,
  // now you have two problems
  const re = /(((https:(?:\/\/)?)(?:[-;:&=+$,\w]+@)?[A-Za-z0-9.-]+|(?:www\.|[-;:&=+$,\w]+@)[A-Za-z0-9.-]+)((?:\/[+~%/.\w\-_]*)?\??(?:[-+=&;%@.\w_]*)#?(?:[.!/\\\w]*))?)/g;
  return markdown.match(re);
};

const find_duplicates = (arr) => {
  const hm = {};
  const dup = [];
  arr.forEach((e) => {
    if (hm[e]) dup.push(e);
    else hm[e] = null;
  });
  return dup;
};

const partition = (arr, func) => {
  const ap = [[], []];
  arr.forEach((e) => (func(e) ? ap[0].push(e) : ap[1].push(e)));
  return ap;
};

async function fetch_link(url) {
  try {
    const { ok, statusText, redirected } = await fetch(url, LINKS_OPTIONS);
    return [url, { ok, status: statusText, redirected }];
  } catch (error) {
    return [url, { ok: false, status: error.message }];
  }
}

async function batch_fetch({ arr, get, post_filter_func, BATCH_SIZE = 8 }) {
  const result = [];
  /* eslint-disable no-await-in-loop */
  for (let i = 0; i < arr.length; i += BATCH_SIZE) {
    const batch = arr.slice(i, i + BATCH_SIZE);
    LOG.debug({ batch });
    let res = await Promise.all(batch.map(get));
    LOG.debug('batch fetched...');
    res = post_filter_func ? res.filter(post_filter_func) : res;
    LOG.debug_string({ res });
    result.push(...res);
  }
  return result;
}

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

// =============================================================
// const batch_github_repos = async (github_links) => {
//   const BATCH_SIZE = 50;
//   const repos = extract_repos(github_links);
//   for (let i = 0; i < repos.length; i += BATCH_SIZE) {
//     const batch = repos.slice(i, i + BATCH_SIZE);
//     const query = generate_GQL_query(batch);
//     LOG.debug({ query });
//     const gql_response = await fetch(
//       'https://api.github.com/graphql',
//       make_GQL_options(query),
//     ).then((r) => r.json());
//     LOG.debug({ gql_response });
//   }
// };
// =============================================================

async function main() {
  const markdown = await fs.readFile(README, 'utf8');
  const links = extract_all_links(markdown);
  const duplicates = find_duplicates(links);
  if (duplicates.length > 0) {
    LOG.error('duplicates', duplicates);
  }
  const [github_links, other_links] = partition(links, (link) =>
    link.startsWith('https://github.com'),
  );

  const other_links_error = await batch_fetch({
    arr: other_links,
    get: fetch_link,
    post_filter_func: (x) => !x[1].ok,
    BATCH_SIZE: 8,
  });
  if (other_links_error.length > 0) {
    LOG.error('other_links_error', other_links_error);
  }

  const repos = extract_repos(github_links);
  const query = generate_GQL_query(repos);
  const options = make_GQL_options(query);
  const gql_response = await fetch(GITHUB_GQL_API, options).then((r) =>
    r.json(),
  );

  const { data } = gql_response;
  if (gql_response.errors) {
    LOG.error_string({ errors: gql_response.errors });
  }
  const repos_fetched = Object.entries(data)
    .map(([, /* k , */ v]) => v.nameWithOwner)
    .sort((a, b) => b - a);

  console.log({ repos_fetched: repos_fetched.length });
}

console.log('starting...');
main();
