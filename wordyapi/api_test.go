package wordyapi

import (
	"encoding/json"
	"testing"
)

func TestParseText(t *testing.T) {

	// Send a string with just two words
	s := TextToParse{Title: "Sample", Text: "test test test"}

	out := ParseText(s)

	var v map[string]interface{}

	err := json.Unmarshal(out, &v)
	if err != nil {
		t.Error("Got this error:", err)
	}

	// test to make sure it counted number of occurances
	if v["test"].(float64) != 3 {
		t.Error("Expected 3 occurances, but got:", v)
	}
}
