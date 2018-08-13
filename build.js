const fs = require('fs-extra');
const fetch = require('node-fetch');
const cheerio = require('cheerio');
const dayjs = require('dayjs');
const showdown = require('showdown');
const Parcel = require('parcel-bundler');
const sm = require('sitemap');

process.env.NODE_ENV = 'production';

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

// --- FILES
const README = 'README.md';
const WEBSITE_FOLDER = 'website';
const DATA_FOLDER = 'data';
const LATEST_FILENAME = `${DATA_FOLDER}/latest`;
const MAPPING = `${DATA_FOLDER}/mapping.json`;
const CATEGORY = `${DATA_FOLDER}/category.json`;
const indexTemplate = `${WEBSITE_FOLDER}/index.tmpl.html`;
const indexDestination = `${WEBSITE_FOLDER}/index.html`;
const tableTemplate = `${WEBSITE_FOLDER}/table.tmpl.html`;
const tableDestination = `${WEBSITE_FOLDER}/table.html`;

// --- CONFIG
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

const sitemapOpts = {
  hostname: 'https://awesome-docker.netlify.com/',
  cacheTime: 6000000, // 600 sec (10 min) cache purge period
  urls: [
    {
      url: '/',
      changefreq: 'daily',
      priority: 0.8,
      lastmodrealtime: true,
      lastmodfile: 'dist/index.html',
    },
    {
      url: '/table.html',
      changefreq: 'daily',
      priority: 0.8,
      lastmodrealtime: true,
      lastmodfile: 'dist/table.html',
    },
  ],
};

// --- FORMAT
const loadEmoji = () =>
  fetch('https://api.github.com/emojis')
    .then(r => r.json())
    .catch(handleFailure);

let emojiMapURL = {};

const emojify = text => {
  if (!text) return text;
  const colonWrapped = /(:[\w\-+]+:)/g;
  const result = text.replace(colonWrapped, match => {
    const name = match.replace(/:/g, '');
    const url = emojiMapURL[name];
    return url ? `<img src="${url}" class="emoji" alt="${name}" />` : match;
  });
  return result || text;
};

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
    categoryName,
  },
  i,
) =>
  [
    `<li data-id="${i}">`,
    `<a href="${repoURL}" class="link ${valueNames[0]}">${name}</a>`,
    `<p class="${valueNames[1]}">${emojify(description) || '-'}</p>`,
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
    } timestamp" title="Stars on GitHub" data-stars="${stargazers}">⭐️${stargazers}</p>`,
    (language && `<p class="${valueNames[5]}">${language}</p>`) || '<p></p>',
    (license &&
      license.url !== null &&
      `<a href="${license.url}" class="link ${valueNames[6]}">${mapLicense(
        license.name,
      )}</a>`) ||
      '<p></p>',
    `<p title="Category">${categoryName}</p>`,
    owner &&
      `<a href="${owner.html_url}" class="link ${valueNames[7]}">${
        owner.login
      }</a>`,
    '</li>',
  ].join('');

const buttonHTLM = valueNames
  .filter(x => !['description', 'homepage'].includes(x))
  .map(v => `<button class="sort" data-sort="${v}">${v} </button>`)
  .join('');

const processMetadata = metaData =>
  [
    `<div class="container">`,
    `<div class="searchbar"><input class="search" placeholder="Search" /></div>`,
    `<div class="sortbtn"><p>Sort by</p>${buttonHTLM}</div>`,
    `</div>`,
    '<ul class="list">',
    Object.values(metaData)
      .map(formatEntry)
      .join(''),
    '</ul>',
  ].join('');

const normalizedMetadata = ([mapping, category, data]) =>
  data.reduce((acc, repo) => {
    const m = mapping[repo.html_url];
    if (!m) {
      console.log('MISSING:', { repo: repo.html_url });
      return acc;
    }
    const c = m && category[m.category];
    if (!c) {
      console.log('CATEGORY MISSING', { mapping: m });
      return acc;
    }
    return {
      ...acc,
      ...{
        [repo.html_url.toLowerCase()]: {
          ...repo,
          ownerType: repo.owner && repo.owner.type,
          categoryName: c.name,
          categoryDescription: c.description,
          status: m.status,
        },
      },
    };
  }, {});

async function processTable() {
  try {
    LOG.debug('Loading files...', { LATEST_FILENAME, tableTemplate });
    const latestFilename = await fs.readFile(LATEST_FILENAME, 'utf8');
    LOG.debug({ latestFilename });

    const data = await Promise.all([
      fs.readJSON(MAPPING),
      fs.readJSON(CATEGORY),
      fs.readJSON(latestFilename),
    ]);

    const metaData = normalizedMetadata(data);
    LOG.debug({ metaData });
    const template = await fs.readFile(tableTemplate, 'utf8');
    LOG.debug('Processing template');
    const $ = cheerio.load(template);
    $('#md').append(processMetadata(metaData));
    LOG.debug('Writing table.html');
    await fs.outputFile(tableDestination, $.html(), 'utf8');
    LOG.debug('✅  DONE 👍');
  } catch (err) {
    handleFailure(err);
  }
}

async function processIndex() {
  const converter = new showdown.Converter({
    omitExtraWLInCodeBlocks: true,
    simplifiedAutoLink: true,
    excludeTrailingPunctuationFromURLs: true,
    literalMidWordUnderscores: true,
    strikethrough: true,
    tables: true,
    tablesHeaderId: true,
    ghCodeBlocks: true,
    tasklists: true,
    disableForced4SpacesIndentedSublists: true,
    simpleLineBreaks: true,
    requireSpaceBeforeHeadingText: true,
    ghCompatibleHeaderId: true,
    ghMentions: true,
    backslashEscapesHTMLTags: true,
    emoji: true,
    splitAdjacentBlockquotes: true,
  });
  // converter.setFlavor('github');

  try {
    LOG.debug('Loading files...', { indexTemplate, README });
    const template = await fs.readFile(indexTemplate, 'utf8');
    const markdown = await fs.readFile(README, 'utf8');

    LOG.debug('Merging files...');
    const $ = cheerio.load(template);
    $('#md').append(converter.makeHtml(markdown));

    LOG.debug('Writing index.html');
    await fs.outputFile(indexDestination, $.html(), 'utf8');
    LOG.debug('DONE 👍');
  } catch (err) {
    handleFailure(err);
  }
}

const bundle = () => {
  LOG.debug('---');
  LOG.debug('📦  Bundling with Parcel.js');
  LOG.debug('---');

  new Parcel(indexDestination, {
    name: 'build',
    publicURL: '/',
  })
    .bundle()
    .then(() =>
      // Creates a sitemap object given the input configuration with URLs
      fs.outputFile(
        'dist/sitemap.xml',
        sm.createSitemap(sitemapOpts).toString(),
      ),
    );
};

async function main() {
  emojiMapURL = await loadEmoji();
  await processTable();
  await processIndex();
  await bundle();
}

main();
