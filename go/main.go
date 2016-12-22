package main

import "github.com/chasestarr/scrabble-cheat/go/dictionary"
import "github.com/chasestarr/scrabble-cheat/go/tree"
import "bytes"
import "fmt"

func getPossibleWords(
	t tree.Tree,
	rack []byte,
	answer []string,
	remaining []byte,
	word []byte) []string {
	if len(remaining) == 0 || !bytes.Contains(remaining, []byte("$")) {
		fmt.Println(word)
		word = word[:len(word)-1]
		answer = append(answer, string(word))
	} else {
		for _, letter := range t.Children {
			fmt.Println(letter)
			ind := bytes.Index(remaining, letter.Value)
			if ind >= 0 {
				tempWord := make([]byte, len(word))
				copy(tempWord, word)
				tempWord = append(tempWord, remaining[ind])

				tempRemaining := make([]byte, len(remaining))
				copy(tempRemaining, remaining)
				tempRemaining = append(tempRemaining[:ind], tempRemaining[ind+1:]...)
				getPossibleWords(letter, rack, answer, tempRemaining, tempWord)
			}
		}
	}

	return answer
}

func main() {
	dictionaryTree := dictionary.GenerateTree()
	rack := "lettersinrack"
	remaining := (rack + "$")
	r := []byte(remaining)
	results := getPossibleWords(dictionaryTree, []byte(rack), []string{}, r, []byte{})

	for _, w := range results {
		fmt.Println(w)
	}
}
