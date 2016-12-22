package dictionary

import (
	"bytes"
	"io/ioutil"
	"log"
)

// ReadDictionary reads scabble dictionary.txt and returns byte slice
func ReadDictionary() [][]byte {
	b, err := ioutil.ReadFile("dictionary/dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}

	e := bytes.Split(b, []byte("\n"))
	return e
}
