const fs = require('fs');
const showdown = require('showdown');
const cheerio = require('cheerio');
const Parcel = require('parcel-bundler');

process.env.NODE_ENV = 'production';

process.on('unhandledRejection', error => {
  console.log('unhandledRejection', error.message);
});

const readme = 'README.md';
const template = 'website/index.tmpl.html';
const merged = 'website/index.html';
const destination = 'website/index.html';

const includeReadme = ({
  md = readme,
  templateHTML = template,
  dest = merged,
} = {}) => {
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

  console.log('Loading files...');
  const indexTemplate = fs.readFileSync(templateHTML, 'utf8');
  const markdown = fs.readFileSync(md, 'utf8');

  console.log('Merging files...');
  const $ = cheerio.load(indexTemplate);
  $('#md').append(converter.makeHtml(markdown));
  console.log('Writing index.html');
  fs.writeFileSync(dest, $.html(), 'utf8');
  console.log('DONE 👍');
};

const bundle = (dest = destination) => {
  console.log('');
  console.log('Bundling with Parcel.js');
  console.log('');

  new Parcel(dest, {
    name: 'build',
    publicURL: '/',
  })
    .bundle()
    .then(() => {
      fs.copyFileSync('website/sitemap.xml', 'dist/sitemap.xml');
    });
};

const main = () => {
  includeReadme();
  bundle();
};

main();
