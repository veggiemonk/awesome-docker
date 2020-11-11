const fs = require('fs-extra');
const helper = require('./common');

console.log({
    DEBUG: process.env.DEBUG || false,
});

const README = 'README.md';

async function main() {
    const has_error = {
        show: false,
        duplicates: '',
        other_links_error: '',
    };
    const markdown = await fs.readFile(README, 'utf8');
    let links = helper.extract_all_links(markdown);
    links = links.filter((l) => !helper.exclude_from_list(l)); // exclude websites
    helper.LOG.debug_string({ links });

    console.log(`total links to check ${links.length}`);

    console.log('checking for duplicates links...');

    const duplicates = helper.find_duplicates(links);
    if (duplicates.length > 0) {
        has_error.show = true;
        has_error.duplicates = duplicates;
    }
    helper.LOG.debug_string({ duplicates });
    const [github_links, external_links] = helper.partition(links, (link) =>
        link.startsWith('https://github.com'),
    );

    console.log(`checking ${external_links.length} external links...`);

    const external_links_error = await helper.batch_fetch({
        arr: external_links,
        get: helper.fetch_link,
        post_filter_func: (x) => !x[1].ok,
        BATCH_SIZE: 8,
    });
    if (external_links_error.length > 0) {
        has_error.show = true;
        has_error.other_links_error = external_links_error;
    }

    console.log(`checking ${github_links.length} GitHub repositories...`);

    console.log(
        `skipping GitHub repository check. Run "npm run test" to execute them manually.`,
    );

    console.log({
        TEST_PASSED: !has_error.show,
        EXTERNAL_LINKS: external_links.length,
    });

    if (has_error.show) {
        helper.LOG.error_string(has_error);
        process.exit(1);
    }
}

console.log('starting...');
main();
