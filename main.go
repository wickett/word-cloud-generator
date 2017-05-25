package main

import (
	"fmt"

	"github.com/wickett/word-cloud-generator/wordyapi"
)

func main() {

	x := wordyapi.TextToParse{"hello", "how are you"}

	fmt.Println(wordyapi.ParseText(x))
}
