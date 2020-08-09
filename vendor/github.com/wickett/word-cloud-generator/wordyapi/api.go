package wordyapi

import (
	"encoding/json"
	"log"
	"strings"
)

type TextToParse struct {
	Title string
	Text  string
}

func ParseText(input TextToParse) []byte {

	words := strings.Split(strings.ToLower(input.Text), " ")

	m := map[string]int{}
	for _, word := range words {
		m[word] = m[word] + 1
	}

	weightJson, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return weightJson

}
