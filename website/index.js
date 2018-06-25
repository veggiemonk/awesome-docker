const List = require('list.js');

const main = () => {
  const userList = new List('md', {
    valueNames: [
      'name',
      'description',
      'homepage',
      'star',
      { name: 'updated', attr: 'data-timestamp' },
      'language',
      'license',
      'author',
    ],
  });
  console.log(`There are ${userList.size()} projects`);
};

main();
