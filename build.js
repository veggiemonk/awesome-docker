const fs = require('fs-extra');
const cheerio = require('cheerio');
const showdown = require('showdown');

process.env.NODE_ENV = 'production';

const LOG = {
  error: (...args) => console.error('❌ ERROR', { ...args }),
  debug: (...args) => {
    if (process.env.DEBUG) console.log('💡 DEBUG: ', { ...args });
  },
};
const handleFailure = (err) => {
  LOG.error(err);
  process.exit(1);
};

process.on('unhandledRejection', handleFailure);

// --- FILES
const README = 'README.md';
const WEBSITE_FOLDER = 'website';
const indexTemplate = `${WEBSITE_FOLDER}/index.tmpl.html`;
const indexDestination = `${WEBSITE_FOLDER}/index.html`;

async function processIndex() {
  const converter = new showdown.Converter();
  converter.setFlavor('github');

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

async function main() {
  await processIndex();
}

main();
