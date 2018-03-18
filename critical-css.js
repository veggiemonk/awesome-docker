const path = require('path');
const fs = require('fs');
const tmpDir = require('os').tmpdir();
const request = require('request');
const criticalcss = require('criticalcss');

const cssUrl =
  'https://awesome-docker.netlify.com/16dc205b0ca3044a54bfb5fc8384de31.css';
const cssPath = path.join(tmpDir, 'style.css');
request(cssUrl)
  .pipe(fs.createWriteStream(cssPath))
  .on('close', () => {
    criticalcss.getRules(cssPath, (err, output) => {
      if (err) {
        throw new Error(err);
      } else {
        criticalcss.findCritical(
          'https://awesome-docker.netlify.com/',
          { rules: JSON.parse(output) },
          (err, output) => {
            if (err) {
              throw new Error(err);
            } else {
              console.log(output);
            }
          }
        );
      }
    });
  });
