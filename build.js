const fs = require('fs');
const showdown = require('showdown');
const cheerio = require('cheerio');
const Parcel = require('parcel-bundler');

process.env.NODE_ENV = 'production';

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
  splitAdjacentBlockquotes: true
});
// converter.setFlavor('github');

console.log('Loading files...');
const index = fs.readFileSync('index.tmpl', 'utf8');
const readme = fs.readFileSync('README.md', 'utf8');

console.log('Merging files...');
const $ = cheerio.load(index);
$('#md').append(converter.makeHtml(readme));

console.log('Writing index.html');
fs.writeFileSync('index.html', $.html(), 'utf8');

console.log('Bundling with Parcel.js');
console.log('');

new Parcel('index.html', {
  name: 'build',
  publicURL: '/awesome-docker'
}).bundle();
