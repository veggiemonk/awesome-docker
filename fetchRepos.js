const fs = require('fs');
const fetch = require('node-fetch');

process.on('unhandledRejection', error => {
  console.log('unhandledRejection', error.message);
});

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
    .then(r => r.json());
}

function delay(ms) {
  return new Promise(resolve => {
    setTimeout(() => resolve(), ms);
  });
}

function extractAllRepos() {
  const re = /https:\/\/github\.com\/([a-zA-Z0-9-._]+)\/([a-zA-Z0-9-._]+)/g;
  const markdown = fs.readFileSync(readme, 'utf-8');
  const md = markdown.match(re);
  const uniq = [...new Set(md)];
  const repos = uniq.map(x => x.slice('https://github.com/'.length, x.length));
  fs.writeFileSync('data/list_repos.json', JSON.stringify(uniq, null, 2));
  return repos;
}

async function main() {
  const repos = extractAllRepos();
  let data = [];
  const batchSize = 10;
  /* eslint-disable no-await-in-loop */
  for (let i = 0; i < repos.length; i += batchSize) {
    const batch = repos.slice(i, i + batchSize);
    console.log({ batch });
    const res = await Promise.all(batch.map(async path => get(path)));
    data = data.concat(res);
    await delay(3000);
  }
  if (process.env.DEBUG) console.log({ data });
  fs.writeFileSync('data/fetched_repo_data.txt', JSON.stringify(data, null, 2));
}

main();
