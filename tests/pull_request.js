const fs = require('fs-extra');
const fetch = require('node-fetch');
const exclude = require('./exclude_in_test.json');

console.log({
  DEBUG: process.env.DEBUG || false,
});

const README = 'README.md';

const LINKS_OPTIONS = {
  redirect: 'error',
  headers: {
    'Content-Type': 'application/json',
    'user-agent':
      'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36',
  },
};

const LOG = {
  error: (...args) => console.error('❌ ERROR', args),
  error_string: (...args) =>
    console.error('❌ ERROR', JSON.stringify({ ...args }, null, '  ')),
  debug: (...args) => {
    if (process.env.DEBUG) console.log('>>> DEBUG: ', { ...args });
  },
  debug_string: (...args) => {
    if (process.env.DEBUG)
      console.log('>>> DEBUG: ', JSON.stringify({ ...args }, null, '  '));
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
  // TODO: replace this mess with a mardown parser ?
  const re = /(((https:(?:\/\/)?)(?:[-;:&=+$,\w]+@)?[A-Za-z0-9.-]+|(?:www\.|[-;:&=+$,\w]+@)[A-Za-z0-9.-]+)((?:\/[+~%/.\w\-_]*)?\??(?:[-+=&;%@.\w_]*)#?(?:[.!/@\-\\\w]*))?)/g;
  return markdown.match(re);
};

const find_duplicates = (arr) => {
  const hm = {};
  const dup = [];
  arr.forEach((e) => {
    if (hm[e]) dup.push(e);
    else hm[e] = true;
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
    LOG.debug_string({ batch });
    let res = await Promise.all(batch.map(get));
    console.log(`batch fetched...${i + BATCH_SIZE}`);
    res = post_filter_func ? res.filter(post_filter_func) : res;
    LOG.debug_string({ res });
    result.push(...res);
  }
  return result;
}

const exclude_length = exclude.length;
const exclude_from_list = (link) => {
  let is_excluded = false;
  for (let i = 0; i < exclude_length; i += 1) {
    if (link.startsWith(exclude[i])) {
      is_excluded = true;
      break;
    }
  }
  return is_excluded;
};

async function main() {
  const has_error = {
    show: false,
    duplicates: '',
    other_links_error: '',
  };
  const markdown = await fs.readFile(README, 'utf8');
  let links = extract_all_links(markdown);
  links = links.filter((l) => !exclude_from_list(l)); // exclude websites
  LOG.debug_string({ links });

  console.log(`total links to check ${links.length}`);

  console.log('checking for duplicates links...');

  const duplicates = find_duplicates(links);
  if (duplicates.length > 0) {
    has_error.show = true;
    has_error.duplicates = duplicates;
  }
  LOG.debug_string({ duplicates });
  const [github_links, external_links] = partition(links, (link) =>
    link.startsWith('https://github.com'),
  );

  console.log(`checking ${external_links.length} external links...`);

  const external_links_error = await batch_fetch({
    arr: external_links,
    get: fetch_link,
    post_filter_func: (x) => !x[1].ok,
    BATCH_SIZE: 8,
  });
  if (external_links_error.length > 0) {
    has_error.show = true;
    has_error.other_links_error = external_links_error;
  }

  console.log(`checking ${github_links.length} GitHub repositories...`);

  console.log(`skipping GitHub repository check. Run "npm run test" to execute them manually.`);

  console.log({
    TEST_PASSED: !has_error.show,
    EXTERNAL_LINKS: external_links.length,
  });

  if (has_error.show) {
    LOG.error_string(has_error);
    process.exit(1);
  }
}

console.log('starting...');
main();
