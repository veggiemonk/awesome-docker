const fs = require('fs');
// const showdown = require('showdown');
const cheerio = require('cheerio');
// const Parcel = require('parcel-bundler');
const dayjs = require('dayjs');
const metaData = require('./data/2018-06-06T17.54.30-fetched_repo_data.json');

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
    (homepage &&
      `<a href="${homepage}" class="link ${valueNames[2]}">🔗 website</a>`) ||
      '<p></p>',
    `<p class="${
      valueNames[3]
    } timestamp" data-timestamp="${stargazers}">⭐️${stargazers}</p>`,
    `<p class="${
      valueNames[4]
    } timestamp" data-timestamp="${updated}">${getLastUpdate(updated)}</p>`,
    (language && `<p class="${valueNames[5]}">💻${language}</p>`) || '<p></p>',
    (license &&
      `<a href="${license.url}" class="link ${valueNames[6]}">📃${
        license.name
      }</a>`) ||
      '<p></p>',
    owner &&
      `<a href="${owner.html_url}" class="link ${valueNames[7]}">Made by ${
        owner.login
      }</a>`,
    '</li>',
  ].join('');

function main() {
  const indexTemplate = fs.readFileSync(templateHTML, 'utf8');
  const $ = cheerio.load(indexTemplate);
  const btn = valueNames.map(
    v => `<button class="sort" data-sort="${v}">Sort by ${v} </button>`,
  );
  $('#md').append(
    [
      `<input class="search" placeholder="Search" /> ${btn.join('')}`,
      '<ul class="list">',
      metaData.map(formatEntry).join(''),
      '</ul>',
    ].join(''),
  );

  $('body').append(
    '<script src="//cdnjs.cloudflare.com/ajax/libs/list.js/1.5.0/list.js"></script>',
  );

  $('body').append(`
  <script> const userList = new List('md', { valueNames:  ['name','description','homepage','star','updated','language','license','author']}); </script>
  `);
  console.log('Writing table.html');
  fs.writeFileSync(destination, $.html(), 'utf8');
  console.log('DONE 👍');
}

main();
