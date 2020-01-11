const fs = require('fs-extra');
const cheerio = require('cheerio');
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
const indexTemplate = `${WEBSITE_FOLDER}/index.tmpl.html`;
const indexDestination = `${WEBSITE_FOLDER}/index.html`;

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
  ],
};

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
  await processIndex();
  await bundle();
}

main();
