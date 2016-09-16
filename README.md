#Scrabble Cheater

This app searches through a sorted tree of English words and matches words that can be produced with a given 'rack' of letters.

##Getting Started

After you clone the repository, install `word-list` but running

    npm install word-list

From there, run

    node src/generateTree.js

to convert the word list into a searchable tree. Note that the tree file takes up approximately 14MB of space.

To search for words, run the `getPossibleWords` function in `src/scrabbleCheat.js` with `dictionaryTree` and a string of your letters as arguments.