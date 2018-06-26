const fs = require('fs');
const cheerio = require('cheerio');
const dayjs = require('dayjs');

const getLatestFilename = fs.readFileSync('data/latest', 'utf-8');
console.log(getLatestFilename);
const metaData = require(`./${getLatestFilename}`); // eslint-disable-line import/no-dynamic-require

process.env.NODE_ENV = 'production';

process.on('unhandledRejection', error => {
  console.log('unhandledRejection', error.message);
});

const templateHTML = 'website/table.tmpl.html';
const destination = 'website/table.html';

const valueNames = [
  'name',
  'description',
  'homepage',
  'star',
  'updated',
  'language',
  'license',
  'author',
];

const getLastUpdate = updated => {
  const updt = Number(dayjs(updated).diff(dayjs(), 'days'));
  if (updt < 0) {
    if (Math.abs(updt) === 1) return `1 day ago`;
    return `${Math.abs(updt)} days ago`;
  } else if (updt === 0) return 'today';
  return updated;
};

const mapHomePage = h => {
  if (h === 'manageiq.org') return 'https://manageiq.org';
  else if (h === 'dev-sec.io') return 'https://dev-sec.io';
  return h;
};

const mapLicense = l => {
  if (l === 'GNU Lesser General Public License v3.0') return 'GNU LGPL v3.0';
  else if (l === 'GNU General Public License v2.0') return 'GNU GPL v2.0';
  else if (l === 'GNU General Public License v3.0') return 'GNU GPL v3.0';
  else if (l === 'BSD 3-Clause "New" or "Revised" License')
    return 'BSD 3-Clause';
  else if (l === 'BSD 2-Clause "Simplified" License') return 'BSD 2-Clause';
  return l;
};

const formatEntry = (
  {
    name,
    html_url: repoURL,
    description,
    homepage,
    stargazers_count: stargazers,
    pushed_at: updated,
    language,
    license,
    owner,
  },
  i,
) =>
  [
    `<li data-id="${i}">`,
    `<a href="${repoURL}" class="link ${valueNames[0]}">${name}</a>`,
    `<p class="${valueNames[1]}">${description || '-'}</p>`,
    `<p class="${
      valueNames[4]
    } timestamp" data-timestamp="${updated}">Last code update: ${getLastUpdate(
      updated,
    )}</p>`,
    (homepage &&
      `<a href="${mapHomePage(homepage)}" class="link ${
        valueNames[2]
      }">website</a>`) ||
      '<p></p>',
    `<p class="${
      valueNames[3]
    } timestamp" data-timestamp="${stargazers}">⭐️${stargazers}</p>`,
    (language && `<p class="${valueNames[5]}">💻${language}</p>`) || '<p></p>',
    (license &&
      license.url !== null &&
      `<a href="${license.url}" class="link ${valueNames[6]}">${mapLicense(
        license.name,
      )}</a>`) ||
      '<p></p>',
    owner &&
      `<p>Made by </p><a href="${owner.html_url}" class="link ${
        valueNames[7]
      }">${owner.login}</a>`,
    '</li>',
  ].join('');

function main() {
  const indexTemplate = fs.readFileSync(templateHTML, 'utf8');
  const $ = cheerio.load(indexTemplate);
  const btn = valueNames.map(
    v => `<button class="sort" data-sort="${v}">${v} </button>`,
  );
  $('#md').append(
    [
      `<div class="container">`,
      `<div class="searchbar" ><input class="search" placeholder="Search" /></div>`,
      `<div class="sortbtn" ><p>Sort by</p>${btn.join('')}</div>`,
      `</div>`,
      '<ul class="list">',
      metaData.map(formatEntry).join(''),
      '</ul>',
    ].join(''),
  );

  console.log('Writing table.html');
  fs.writeFileSync(destination, $.html(), 'utf8');
  console.log('DONE 👍');
}

main();
