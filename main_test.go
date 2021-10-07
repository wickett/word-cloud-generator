package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSimpleSentence(t *testing.T) {
	result := testFixture(t, []byte(`{"Text":"TEST this is only a test"}`))

	// Check the response body is what we expect.
	expected := `{ "a": 1, "is": 1, "only": 1, "test": 2, "this": 1 }`

	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}

func TestLongPunc(t *testing.T) {
	result := testFixture(t, []byte(`{"Text":"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sem integer vitae justo eget magna fermentum iaculis. Adipiscing at in tellus integer feugiat scelerisque varius morbi enim. Elementum sagittis vitae et leo duis ut. Posuere lorem ipsum dolor sit. Arcu dictum varius duis at. Ridiculus mus mauris vitae ultricies leo integer malesuada. Arcu non sodales neque sodales ut. Diam sollicitudin tempor id eu nisl nunc mi ipsum. Ultricies mi eget mauris pharetra et ultrices neque ornare aenean. Feugiat scelerisque varius morbi enim nunc faucibus a pellentesque. Morbi leo urna molestie at. Sit amet nisl purus in. Nisl tincidunt eget nullam non."}`))

	// Check the response body is what we expect.
	expected := `{ "a": 1, "adipiscing": 2, "aenean.": 1, "aliqua.": 1, "amet": 1, "amet,": 1, "arcu": 2, "at": 1, "at.": 2, "consectetur": 1, "diam": 1, "dictum": 1, "do": 1, "dolor": 2, "dolore": 1, "duis": 2, "eget": 3, "eiusmod": 1, "elementum": 1, "elit,": 1, "enim": 1, "enim.": 1, "et": 3, "eu": 1, "faucibus": 1, "fermentum": 1, "feugiat": 2, "iaculis.": 1, "id": 1, "in": 1, "in.": 1, "incididunt": 1, "integer": 3, "ipsum": 2, "ipsum.": 1, "justo": 1, "labore": 1, "leo": 3, "lorem": 2, "magna": 2, "malesuada.": 1, "mauris": 2, "mi": 2, "molestie": 1, "morbi": 3, "mus": 1, "neque": 2, "nisl": 3, "non": 1, "non.": 1, "nullam": 1, "nunc": 2, "ornare": 1, "pellentesque.": 1, "pharetra": 1, "posuere": 1, "purus": 1, "ridiculus": 1, "sagittis": 1, "scelerisque": 2, "sed": 1, "sem": 1, "sit": 2, "sit.": 1, "sodales": 2, "sollicitudin": 1, "tellus": 1, "tempor": 2, "tincidunt": 1, "ultrices": 1, "ultricies": 2, "urna": 1, "ut": 1, "ut.": 2, "varius": 3, "vitae": 3 }`

	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}

func testFixture(t *testing.T, jsonString []byte) string {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	var jsonStr = jsonString
	url := "/api"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(receiveJSONHandler)

	// Our handlers satisfy http.Handler, so we can call the ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	strOut := strings.Join(strings.Fields(rr.Body.String()), " ")

	return strOut
}
