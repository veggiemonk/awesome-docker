const fs = require('fs');
const dayjs = require('dayjs');

function getLastUpdate(updated) {
  const updt = Number(dayjs(updated).diff(dayjs(), 'days'));
  if (updt < 0) {
    if (Math.abs(updt) === 1) return `1 day ago`;
    return `${Math.abs(updt)} days ago`;
  } else if (updt === 0) return 'today';
  return updated;
}

function createLine(data) {
  const {
    name,
    html_url: repoURL,
    description,
    homepage,
    stargazers_count: stargazers,
    updated_at: updated,
    language,
    license,
    owner,
  } = data;
  if (!name) return '|ERROR |';

  const lineData = [
    `[${name}](${repoURL})`,
    description || '-',
    homepage || '-',
    stargazers,
    getLastUpdate(updated),
    language,
    license && `[${license.name}](${license.url})`,
    owner && `[${owner.login}](${owner.html_url})`,
  ];
  return `|${lineData.join('|')}|`;
}

function main() {
  const raw = fs.readFileSync('data/fetched_repo_data.json', 'utf-8');
  const data = JSON.parse(raw);
  const header = `
| Name        | Description | Homepage | Star | Updated | Language | License | Author |
| ----------- | ----------- | -------- | ---- | ------- | -------- | :---:   | ------:|`;
  const table = [header]
    .concat(
      data
        .sort((a, b) => Number(b.stargazers_count) - Number(a.stargazers_count))
        .map(createLine),
    )
    .join('\n');

  fs.writeFileSync('data/table.md', table);
}

main();
