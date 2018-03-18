const fs = require('fs');
const showdown = require('showdown');
const cheerio = require('cheerio');
const Parcel = require('parcel-bundler');
const critical = require('critical');

process.env.NODE_ENV = 'production';

const includeReadme = () => {
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
  return { base: 'website/', src: 'index.merged.html' };
};

const css = ({ base, src }) => {
  console.log('');
  console.log('Generating critical css above the fold');
  console.log('');
  const dimensions = [
    {
      height: 200,
      width: 500
    },
    {
      height: 900,
      width: 1200
    }
  ];
  const options = {
    inline: true,
    base,
    src,
    dest: 'index.html',
    dimensions
  };

  return critical.generate(options);
};

const bundle = () => {
  console.log('');
  console.log('Bundling with Parcel.js');
  console.log('');

  new Parcel('website/index.html', {
    name: 'build',
    publicURL: '/'
  }).bundle();
};

const main = async () => {
  const { base, src } = includeReadme();
  await css({ base, src });
  bundle();
};

main();
