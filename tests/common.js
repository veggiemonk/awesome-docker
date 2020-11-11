const fetch = require('node-fetch');
const exclude = require('./exclude_in_test.json');

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

module.exports = {
  LOG,
  handleFailure,
  extract_all_links,
  find_duplicates,
  partition,
  fetch_link,
  batch_fetch,
  exclude_from_list,
};
