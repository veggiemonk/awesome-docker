const fs = require('fs');
const cheerio = require('cheerio');
const dayjs = require('dayjs');
const icons = require('./icons');

const getLatestFilename = fs.readFileSync('data/latest', 'utf-8');
console.log(getLatestFilename);
const metaData = require(`./${getLatestFilename}`); // eslint-disable-line import/no-dynamic-require

process.env.NODE_ENV = 'production';

process.on('unhandledRejection', error => {
  console.log('unhandledRejection', error.message);
});

// const table = 'data/table.md';
const templateHTML = 'website/table.tmpl.html';
// const merged = 'website/table.html';
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

const formatEntry = (
  {
    name,
    html_url: repoURL,
    description,
    homepage,
    stargazers_count: stargazers,
    updated_at: updated,
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
    } timestamp" data-timestamp="${updated}">Last update: ${getLastUpdate(
      updated,
    )}</p>`,
    (homepage &&
      `<a href="${homepage}" class="link ${valueNames[2]}">🔗 website</a>`) ||
      '<p></p>',
    `<p class="${
      valueNames[3]
    } timestamp" data-timestamp="${stargazers}">⭐️${stargazers}</p>`,
    (language &&
      `<p class="${valueNames[5]}">${icons[language] ||
        '💻'}${language}</p>`) ||
      '<p></p>',
    (license &&
      `<a href="${license.url}" class="link ${valueNames[6]}">📃 ${
        license.name
      }</a>`) ||
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
      // `<ul class="pagination"></ul>`,
      `</div>`,
      '<ul class="list">',
      metaData.map(formatEntry).join(''),
      '</ul>',
    ].join(''),
  );

  $('body').append(
    '<script src="//cdnjs.cloudflare.com/ajax/libs/list.js/1.5.0/list.js"></script>',
  );

  $('body').append(
    [
      '<script>',
      `const userList = new List('md', {`,
      ` valueNames:  ['name','description','homepage','star','updated','language','license','author'],`,
      // ` page: 20,`,
      // ` pagination: true,`,
      `});`,
      '</script>',
    ].join(''),
  );
  console.log('Writing table.html');
  fs.writeFileSync(destination, $.html(), 'utf8');
  console.log('DONE 👍');
}

main();
