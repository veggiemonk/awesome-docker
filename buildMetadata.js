const fs = require('fs-extra');
const fetch = require('node-fetch');
const dayjs = require('dayjs');

require('draftlog').into(console);

const LOG = {
  error: (...args) => console.error('❌ ERROR', { ...args }),
  debug: (...args) => {
    if (process.env.DEBUG) console.log('💡 DEBUG: ', { ...args });
  },
};
const handleFailure = err => {
  LOG.error(err);
  process.exit(1);
};

process.on('unhandledRejection', handleFailure);

if (!process.env.GITHUB_TOKEN) {
  LOG.error('no credentials found.');
  process.exit(1);
}

const TOKEN = process.env.GITHUB_TOKEN

// --- ENV VAR ---
const BATCH_SIZE = parseInt(process.env.BATCH_SIZE, 10) || 10;
const DELAY = parseInt(process.env.DELAY, 10) || 3000;
const INTERVAL = parseInt(process.env.INTERVAL, 10) || 1;
const INTERVAL_UNIT = process.env.INTERVAL_UNIT || 'days';

// --- FILES ---
const DATA_FOLDER = 'data';
const README = 'README.md';
const DATE = dayjs().format('YYYY-MM-DDTHH.mm.ss');
const GITHUB_METADATA_FILE = `${DATA_FOLDER}/${DATE}-fetched_repo_data.json`;
const LATEST_FILENAME = `${DATA_FOLDER}/latest`;
const GITHUB_REPOS = `${DATA_FOLDER}/repository.json`;
const Authorization = `token ${TOKEN}`

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
const removeHost = x => x.slice('https://github.com/'.length, x.length);

const delay = ms =>
  new Promise(resolve => {
    setTimeout(() => resolve(), ms);
  });

const get = (pathURL, opt) => {
  LOG.debug(`Fetching ${pathURL}`);
  return fetch(`${API}repos/${pathURL}`, {
    ...options,
    ...opt,
  })
    .catch(handleFailure)
    .then(response => {
      if (response.ok) return response.json();
      throw new Error('Network response was not ok.');
    })
    .catch(handleFailure);
};

const fetchAll = batch => Promise.all(batch.map(async pathURL => get(pathURL)));

const extractAllRepos = markdown => {
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
    // poor man's rate limiting so github don't ban us
    await delay(DELAY);
  }
  ProgressBar(repos.length, BATCH_SIZE, repos.length);
  return metadata;
}

function shouldUpdate(lastUpdateTime) {
  LOG.debug({ lastUpdateTime });
  if (!lastUpdateTime) return true;

  const hours = lastUpdateTime.slice(
    'data/YYYY-MM-DDT'.length,
    'data/YYYY-MM-DDTHH'.length,
  );
  const latestUpdate = dayjs(
    lastUpdateTime.slice('data/'.length, 'data/YYYY-MM-DD'.length),
  ).add(hours, 'hour');

  LOG.debug({ latestUpdate: latestUpdate.format() });

  return dayjs().diff(latestUpdate, INTERVAL_UNIT) >= INTERVAL;
}

async function main() {
  try {
    const lastUpdateTime = await fs.readFile(LATEST_FILENAME, 'utf8');
    LOG.debug('Checking if updating is needed');
    if (!shouldUpdate(lastUpdateTime)) {
      LOG.debug('Last update was less than a day ago 😅. Exiting...');
      process.exit();
    }

    const markdown = await fs.readFile(README, 'utf8');
    const githubRepos = extractAllRepos(markdown);
    LOG.debug('writing repo list to disk...');
    await fs.outputJSON(GITHUB_REPOS, githubRepos, { spaces: 2 });

    LOG.debug('fetching data...');
    const metadata = await batchFetchRepoMetadata(githubRepos);

    LOG.debug('writing metadata to disk...');
    await fs.outputJSON(GITHUB_METADATA_FILE, metadata, { spaces: 2 });
    LOG.debug('✅ metadata saved');

    LOG.debug('removing latest...');
    await fs.remove(LATEST_FILENAME);

    LOG.debug('writing latest...');
    await fs.outputFile(LATEST_FILENAME, GITHUB_METADATA_FILE);
    LOG.debug('✅ late update time saved', {
      LATEST_FILENAME,
      GITHUB_METADATA_FILE,
    });

    LOG.debug('gracefully shutting down.');
    process.exit();
  } catch (err) {
    handleFailure(err);
  }
}

main();
