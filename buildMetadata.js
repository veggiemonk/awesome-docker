const fs = require('fs-extra');
const fetch = require('node-fetch');
const dayjs = require('dayjs');

require('draftlog').into(console);

process.on('unhandledRejection', error => {
  console.log('unhandledRejection', error.message);
});

if (!process.env.TOKEN) {
  console.log('❌  no credentials found.');
  process.exit(1);
}

// --- ENV VAR ---
const BATCH_SIZE = parseInt(process.env.BATCH_SIZE, 10) || 10;
const DELAY = parseInt(process.env.DELAY, 10) || 3000;
// --- FILENAME ---
const README = 'README.md';
const DATA_FOLDER = 'data';
const GITHUB_METADATA_FILE = `${DATA_FOLDER}/${dayjs().format(
  'YYYY-MM-DDTHH.mm.ss',
)}-fetched_repo_data.json`;
const LATEST_FILENAME = `${DATA_FOLDER}/latest`;
const GITHUB_REPOS = `${DATA_FOLDER}/list_repos.json`;
// --- HTTP ---
const API = 'https://api.github.com/';
const options = {
  method: 'GET',
  headers: {
    'User-Agent': 'awesome-docker script listing',
    'Content-Type': 'application/json',
    Authorization: `token ${process.env.TOKEN}`,
  },
};

const removeHost = x => x.slice('https://github.com/'.length, x.length);
const barLine = console.draft('Starting batch...');
const handleFailure = err => {
  console.error('❌ ERROR', err);
  process.exit(1);
};
const LOG = (...logs) => {
  if (process.env.DEBUG) console.log(...logs);
};
const delay = ms =>
  new Promise(resolve => {
    setTimeout(() => resolve(), ms);
  });

const get = (pathURL, opt) => {
  LOG(` Fetching ${pathURL}`);
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
  return barLine(
    `[${'='.repeat(units)}${' '.repeat(50 - units)}] ${progress}%  -  # ${i}`,
  );
};

async function batchFetchRepoMetadata(githubRepos) {
  const repos = githubRepos.map(removeHost);

  const metadata = [];
  /* eslint-disable no-await-in-loop */
  for (let i = 0; i < repos.length; i += BATCH_SIZE) {
    const batch = repos.slice(i, i + BATCH_SIZE);
    LOG({ batch });
    const res = await fetchAll(batch);
    LOG('batch fetched...');
    metadata.push(...res);
    ProgressBar(i, BATCH_SIZE, repos.length);
    // poor man's rate limiting so github don't ban us
    await delay(DELAY);
  }
  ProgressBar(repos.length, BATCH_SIZE, repos.length);
  return metadata;
}

function shouldUpdate(fileLatestUpdate) {
  LOG({ fileLatestUpdate });
  if (!fileLatestUpdate) return true;

  const hours = fileLatestUpdate.slice(
    'data/YYYY-MM-DDT'.length,
    'data/YYYY-MM-DDTHH'.length,
  );
  const latestUpdate = dayjs(
    fileLatestUpdate.slice('data/'.length, 'data/YYYY-MM-DD'.length),
  ).add(hours, 'hour');
  LOG({ latestUpdate: latestUpdate.format() });
  const isMoreThanOneDay = dayjs().diff(latestUpdate, 'hours') >= 1;
  return isMoreThanOneDay;
}

async function main() {
  try {
    const getLatest = await fs.readFile(LATEST_FILENAME, 'utf8');
    LOG('Checking if updating is needed');
    if (!shouldUpdate(getLatest)) {
      console.log('Last update was less than a day ago 😅. Exiting...');
      process.exit(1);
    }
    const markdown = await fs.readFile(README, 'utf8');
    const githubRepos = extractAllRepos(markdown);
    LOG('writing repo list to disk...');
    await fs.outputJSON(GITHUB_REPOS, githubRepos, { spaces: 2 });

    LOG('fetching data...');
    const metadata = await batchFetchRepoMetadata(githubRepos);

    LOG('writing metadata to disk...');
    await fs.outputJSON(GITHUB_METADATA_FILE, metadata, { spaces: 2 });
    LOG('✅ metadata saved');

    LOG('removing latest...');
    await fs.remove(LATEST_FILENAME);

    LOG('writing latest...');
    await fs.outputFile(LATEST_FILENAME, GITHUB_METADATA_FILE);

    if (process.env.DEBUG) {
      console.log('✅ late update time saved', {
        LATEST_FILENAME,
        GITHUB_METADATA_FILE,
      });
    }

    console.log('gracefully shutting down.');
    process.exit();
  } catch (err) {
    handleFailure(err);
  }
}

main();
