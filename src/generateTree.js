var fs = require('fs');
var wordListPath = require('word-list');
var wordArray = fs.readFileSync(wordListPath, 'utf8').split('\n');
var Tree = require('./tree.js');
var wordTree = Tree();

var addWordToTree = function(tree, word) {
  word = word + '$';
  word = word.split('');
  var currentLeaf = tree;
  for (var i = 0; i < word.length; i++) {
    var index = -1;
    currentLeaf.c.forEach((child, ind) => {
      if (child.v === word[i]) {
        index = ind;
      }
    });
    if (index >= 0) {
      currentLeaf = currentLeaf.c[index];
    } else {
      currentLeaf = currentLeaf.addChild(word[i]);
    }
  }
};

wordArray.forEach((word) => { addWordToTree(wordTree, word)} );

fs.writeFileSync('lib/dictionaryTreeJSON.json', JSON.stringify(wordTree));








