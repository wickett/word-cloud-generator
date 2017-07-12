package wordyapi

import (
	"encoding/json"
	"testing"
)

func TestParseText(t *testing.T) {
	//	t.Error("This test failed")

	// Send a string with mixed case
	s := TextToParse{Title: "Sample", Text: "Test test test Test test test TEST"}
	out := ParseText(s)

	var v map[string]interface{}

	err := json.Unmarshal(out, &v)
	if err != nil {
		t.Error("Got this error:", err)
	}

	// test to make sure it counted the variations of capitalization
	if v["test"].(float64) != 7 {
		t.Error("Expected 7 occurances of test, but got:", v)
	}
}
