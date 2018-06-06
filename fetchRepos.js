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

// --- UTILS ---
function get(path, opt) {
  return fetch(`${API}repos/${path}`, {
    ...options,
    ...opt,
  })
    .catch(err => console.error(err))
    .then(r => {
      if (r.ok) return r.json();
      throw new Error('Network response was not ok.');
    })
    .catch(err => console.error(err));
}
const delay = ms =>
  new Promise(resolve => {
    setTimeout(() => resolve(), ms);
  });

const extractAllRepos = markdown => {
  const re = /https:\/\/github\.com\/([a-zA-Z0-9-._]+)\/([a-zA-Z0-9-._]+)/g;
  const md = markdown.match(re);
  return [...new Set(md)];
};

const barLine = console.draft('Starting batch...');

const ProgressBar = (i, batchSize, total) => {
  const progress = Math.round((i / total) * 100);
  const units = Math.round(progress / 2);
  return barLine(
    `[${'='.repeat(units)}${' '.repeat(50 - units)}] ${progress}%  -  # ${i}`,
  );
};

const removeHost = x => x.slice('https://github.com/'.length, x.length);

const readFile = promisify(fs.readFile);

// ------------------------------------------------------------

async function main() {
  try {
    const markdown = await readFile(README, { encoding: 'utf8' });
    const githubRepos = extractAllRepos(markdown);
    fs.writeFile(
      GITHUB_REPOS,
      JSON.stringify(githubRepos, null, 2),
      err => err && console.error('FILE ERROR', err),
    );

    const repos = githubRepos.map(removeHost);

    /* eslint-disable no-await-in-loop */
    for (let i = 0; i < repos.length; i += BATCH_SIZE) {
      const batch = repos.slice(i, i + BATCH_SIZE);
      if (process.env.DEBUG) console.log({ batch });
      const res = await Promise.all(batch.map(async path => get(path)));
      fs.appendFile(
        GITHUB_METADATA_FILE,
        JSON.stringify(res, null, 2),
        err => err && console.error(err),
      );
      ProgressBar(i, BATCH_SIZE, repos.length);
      await delay(DELAY);
    }
    ProgressBar(repos.length, BATCH_SIZE, repos.length);
  } catch (err) {
    console.error('ERROR:', err);
  }
}

main();
