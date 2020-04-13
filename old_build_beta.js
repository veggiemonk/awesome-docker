const fs = require('fs-extra');
const fetch = require('node-fetch');

require('draftlog').into(console);

const LOG = {
  error: (...args) => console.error('  ERROR', { ...args }),
  debug: (...args) => {
    if (process.env.DEBUG) console.log('💡 DEBUG: ', { ...args });
  },
};
const handleFailure = (err) => {
  LOG.error(err);
  process.exit(1);
};

process.on('unhandledRejection', handleFailure);

if (!process.env.GITHUB_TOKEN) {
  LOG.error('no credentials found.');
  process.exit(1);
}

const TOKEN = process.env.GITHUB_TOKEN;

// --- ENV VAR ---
const BATCH_SIZE = parseInt(process.env.BATCH_SIZE, 10) || 10;
const DELAY = parseInt(process.env.DELAY, 10) || 3000;
const INTERVAL = parseInt(process.env.INTERVAL, 10) || 1;
const INTERVAL_UNIT = process.env.INTERVAL_UNIT || 'days';

// --- FILES ---
const DATA_FOLDER = 'data';
const README = 'README.md';
const LATEST_FILENAME = `${DATA_FOLDER}/latest`;
const GITHUB_REPOS = `${DATA_FOLDER}/repository.json`;
const Authorization = `token ${TOKEN}`;

// --- HTTP ---
const API = 'https://api.github.com/';
const options = {
  method: 'GET',
  headers: {
    'User-Agent': 'awesome-docker script listing',
    'Content-Type': 'application/json',
    Authorization,
  },
};

// ----------------------------------------------------------------------------
const removeHost = (x) => x.slice('https://github.com/'.length, x.length);

const delay = (ms) =>
  new Promise((resolve) => {
    setTimeout(() => resolve(), ms);
  });

const get = (pathURL, opt) => {
  LOG.debug(`Fetching ${pathURL}`);
  return fetch(`${API}repos/${pathURL}`, {
    ...options,
    ...opt,
  })
    .catch(handleFailure)
    .then((response) => {
      if (response.ok) return response.json();
      throw new Error('Network response was not ok.');
    })
    .catch(handleFailure);
};

const fetchAll = (batch) =>
  Promise.all(batch.map(async (pathURL) => get(pathURL)));

const extractAllLinks = (markdown) => {
  const re = /((([A-Za-z]{3,9}:(?:\/\/)?)(?:[\-;:&=\+\$,\w]+@)?[A-Za-z0-9\.\-]+|(?:www\.|[\-;:&=\+\$,\w]+@)[A-Za-z0-9\.\-]+)((?:\/[\+~%\/\.\w\-_]*)?\??(?:[\-\+=&;%@\.\w_]*)#?(?:[\.\!\/\\\w]*))?)/g;
  return markdown.match(re);
};

const extractAllRepos = (markdown) => {
  const re = /https:\/\/github\.com\/([a-zA-Z0-9-._]+)\/([a-zA-Z0-9-._]+)/g;
  const md = markdown.match(re);
  return [...new Set(md)];
};

const ProgressBar = (i, batchSize, total) => {
  const progress = Math.round((i / total) * 100);
  const units = Math.round(progress / 2);
  const barLine = console.draft('Starting batch...');
  return barLine(
    `[${'='.repeat(units)}${' '.repeat(50 - units)}] ${progress}%  -  # ${i}`,
  );
};
// ----------------------------------------------------------------------------
async function batchFetchRepoMetadata(githubRepos) {
  const repos = githubRepos.map(removeHost);

  const metadata = [];
  /* eslint-disable no-await-in-loop */
  for (let i = 0; i < repos.length; i += BATCH_SIZE) {
    const batch = repos.slice(i, i + BATCH_SIZE);
    LOG.debug({ batch });
    const res = await fetchAll(batch);
    LOG.debug('batch fetched...');
    metadata.push(...res);
    ProgressBar(i, BATCH_SIZE, repos.length);
    // poor man's rate limiting so github doesn't ban us
    await delay(DELAY);
  }
  ProgressBar(repos.length, BATCH_SIZE, repos.length);
  return metadata;
}

async function main() {
  try {
    const markdown = await fs.readFile(README, 'utf8');
    const links = extractAllLinks(markdown);
    const githubRepos = extractAllRepos(markdown);
    LOG.debug('writing repo list to disk...');
    await fs.outputJSON(GITHUB_REPOS, githubRepos, { spaces: 2 });

    LOG.debug('fetching data...');
    const metadata = await batchFetchRepoMetadata(githubRepos);

    LOG.debug('gracefully shutting down.');
    process.exit();
  } catch (err) {
    handleFailure(err);
  }
}

main();
