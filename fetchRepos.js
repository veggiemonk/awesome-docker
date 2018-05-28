const fs = require('fs');
const fetch = require('node-fetch');

require('draftlog').into(console);

process.on('unhandledRejection', error => {
  console.log('unhandledRejection', error.message);
});

if (!process.env.TOKEN) {
  throw new Error('no github token found');
}
const BATCH_SIZE = parseInt(process.env.BATCH_SIZE, 10) || 10;
const DELAY = parseInt(process.env.DELAY, 10) || 3000;
const readme = 'README.md';
const API = 'https://api.github.com/';
const options = {
  method: 'GET',
  headers: {
    'User-Agent': 'awesome-docker script listing',
    'Content-Type': 'application/json',
    Authorization: `token ${process.env.TOKEN}`,
  },
};

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

function delay(ms) {
  return new Promise(resolve => {
    setTimeout(() => resolve(), ms);
  });
}

function extractAllRepos(markdown) {
  const re = /https:\/\/github\.com\/([a-zA-Z0-9-._]+)\/([a-zA-Z0-9-._]+)/g;
  const md = markdown.match(re);
  return [...new Set(md)];
}

const barLine = console.draft('Starting batch...');

function ProgressBar(i, batchSize, total) {
  const progress = Math.round(i / total * 100);
  const units = Math.round(progress / 2);
  return barLine(
    `[${'='.repeat(units)}${' '.repeat(50 - units)}] ${progress}%  -  # ${i}`,
  );
}

async function main() {
  const markdown = fs.readFileSync(readme, 'utf-8');
  const githubRepos = extractAllRepos(markdown);
  const repos = githubRepos.map(x =>
    x.slice('https://github.com/'.length, x.length),
  );
  fs.writeFileSync(
    'data/list_repos.json',
    JSON.stringify(githubRepos, null, 2),
  );
  let data = [];

  /* eslint-disable no-await-in-loop */
  for (let i = 0; i < repos.length; i += BATCH_SIZE) {
    const batch = repos.slice(i, i + BATCH_SIZE);
    const res = await Promise.all(batch.map(async path => get(path)));
    data = data.concat(res);
    if (process.env.DEBUG) console.log({ batch });
    ProgressBar(i, BATCH_SIZE, repos.length);
    await delay(DELAY);
  }
  ProgressBar(repos.length, BATCH_SIZE, repos.length);
  if (process.env.DEBUG) console.log({ data });
  fs.writeFileSync(
    'data/fetched_repo_data.json',
    JSON.stringify(data, null, 2),
  );
}

main();
