const fs = require('fs');
const showdown = require('showdown');
const cheerio = require('cheerio');
const Parcel = require('parcel-bundler');
const critical = require('critical');

process.env.NODE_ENV = 'production';

const main = () => {
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
  const index = fs.readFileSync('website/index.tmpl.html', 'utf8');
  const readme = fs.readFileSync('README.md', 'utf8');

  console.log('Merging files...');
  const $ = cheerio.load(index);
  $('#md').append(converter.makeHtml(readme));

  console.log('Writing index.html');
  fs.writeFileSync('website/index.merged.html', $.html(), 'utf8');

  console.log('');
  console.log('Generating critical css above the fold');
  console.log('');

  critical
    .generate({
      inline: true,
      base: 'website/',
      src: 'index.merged.html',
      dest: 'index.html',
      css: 'website/style.css',
      dimensions: [
        {
          height: 200,
          width: 500
        },
        {
          height: 900,
          width: 1200
        }
      ]
    })
    .then(() => {
      console.log('Bundling with Parcel.js');
      console.log('');

      new Parcel('website/index.html', {
        name: 'build',
        // publicURL: '/awesome-docker'
        publicURL: '/'
      }).bundle();
    });
};

main();
