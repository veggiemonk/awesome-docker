const fs = require('fs');
const { promisify } = require('util');
const fetch = require('node-fetch');
const dayjs = require('dayjs');

require('draftlog').into(console);

process.on('unhandledRejection', error => {
  console.log('unhandledRejection', error.message);
});

if (!process.env.TOKEN) {
  throw new Error('no github token found');
}

// --- ENV VAR ---
const BATCH_SIZE = parseInt(process.env.BATCH_SIZE, 10) || 10;
const DELAY = parseInt(process.env.DELAY, 10) || 3000;
// --- FILENAME ---
const README = 'README.md';
const GITHUB_METADATA_FILE = `data/${dayjs().format(
  'YYYY-MM-DDTHH.mm.ss',
)}-fetched_repo_data.json`;
const LATEST_FILENAME = 'data/latest';
const GITHUB_REPOS = 'data/list_repos.json';
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
const readFile = promisify(fs.readFile);
const writeFile = promisify(fs.writeFile);
const printError = err => err && console.error('❌ ERROR', err);
const barLine = console.draft('Starting batch...');

const delay = ms =>
  new Promise(resolve => {
    setTimeout(() => resolve(), ms);
  });

const get = (path, opt) =>
  fetch(`${API}repos/${path}`, {
    ...options,
    ...opt,
  })
    .catch(printError)
    .then(response => {
      if (response.ok) return response.json();
      throw new Error('Network response was not ok.');
    })
    .catch(printError);

const fetchAll = batch => Promise.all(batch.map(async path => get(path)));

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
    if (process.env.DEBUG) console.log({ batch });
    const res = await fetchAll(batch);
    metadata.push(...res);
    ProgressBar(i, BATCH_SIZE, repos.length);
    // poor man's rate limiting so github don't ban us
    await delay(DELAY);
  }
  ProgressBar(repos.length, BATCH_SIZE, repos.length);
  return metadata;
}

async function main() {
  try {
    const markdown = await readFile(README, { encoding: 'utf8' });
    const githubRepos = extractAllRepos(markdown);
    await writeFile(
      GITHUB_REPOS,
      JSON.stringify(githubRepos, null, 2),
      printError,
    );

    const metadata = await batchFetchRepoMetadata(githubRepos);

    await writeFile(
      GITHUB_METADATA_FILE,
      JSON.stringify(metadata, null, 2),
      printError,
    );
    console.log('✅ metadata saved');

    // save the latest
    fs.writeFile(LATEST_FILENAME, GITHUB_METADATA_FILE, printError);
  } catch (err) {
    printError(err);
  }
}

main();
