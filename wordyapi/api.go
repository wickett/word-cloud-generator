package wordyapi

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/wickett/wordy/ignorewords"
)

type TextToParse struct {
	Title string
	Text  string
}

func ParseText(input TextToParse) []byte {

	words := strings.Split(strings.ToLower(input.Text), " ")

	//fmt.Printf("%q\n", words)
	m := map[string]int{}
	for _, word := range words {
		if ignorewords.IsUseless(word) {
			delete(m, word)
		}
		m[word] = m[word] + 1
	}

	weightJson, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return weightJson

}
