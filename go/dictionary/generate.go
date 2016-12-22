package dictionary

import (
	"bytes"

	"github.com/chasestarr/scrabble-cheat/go/tree"
)

func addWordToTree(t *tree.Tree, word []byte) {
	w := bytes.Split(word, []byte(""))
	w = append(w, []byte("$"))
	currentLeaf := t
	for _, c := range w {
		index := -1
		for ind, child := range currentLeaf.Children {
			if bytes.Compare(child.Value, c) == 0 {
				index = ind
			}
		}

		if index >= 0 {
			currentLeaf = &currentLeaf.Children[index]
		} else {
			currentLeaf = currentLeaf.Add(c)
		}
	}
}

// GenerateTree returns scrabble word trie
func GenerateTree() tree.Tree {
	e := ReadDictionary()
	wordTree := tree.Tree{}

	// populate tree
	for _, c := range e {
		addWordToTree(&wordTree, c)
	}

	return wordTree
}
