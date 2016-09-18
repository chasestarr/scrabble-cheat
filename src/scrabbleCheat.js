var fs = require('fs');
var dictionaryTree = fs.readFileSync('lib/dictionaryTreeJSON.json');
dictionaryTree = JSON.parse(dictionaryTree);

var getPossibleWords = function(tree, rack) {
  var answers = arguments[2] || [];
  var remainingRack = arguments[3] || (rack + '$').split('');
  var builtWord = arguments[4] || [];

  if (remainingRack.length === 0 || remainingRack.indexOf('$') === -1) {
    // rack is empty -- push the word we've built to answers
    builtWord.pop(); // take off the trailing '$'
    answers.push(builtWord.join(''));
  } else {
    // letters are left in rack
    // see if we can make words from our letters
    tree.c.forEach((letter, childIndex) => {
      var ind = remainingRack.indexOf(letter.v);
      if (ind >= 0) {
        var tempBuiltWord = builtWord.slice();
        tempBuiltWord.push(remainingRack[ind]);
        var tempRemainingRack = remainingRack.slice();
        tempRemainingRack.splice(ind,1);
        getPossibleWords(letter, rack, answers, tempRemainingRack, tempBuiltWord);
      }
    });
  }

  return answers;
};

var rack = 'lettersinrack';
var results = getPossibleWords(dictionaryTree, rack);
console.log('\n'.repeat(5));
console.log('/'.repeat(15));
console.log('Found ' + results.length + ' words with rack: ', rack, '.');
console.log('v'.repeat(15));
console.log(results);
console.log('^'.repeat(15));
console.log('Found ' + results.length + ' words with rack: ', rack, '.');
console.log('/////////////////////////////');


