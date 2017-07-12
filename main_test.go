package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReceiveJSONHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	var jsonStr = []byte(`{"Text":"TEST this is only a test"}`)
	url := "/api"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(receiveJSONHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	strOut := strings.Join(strings.Fields(rr.Body.String()), " ")
	// Check the response body is what we expect.
	expected := `{ "a": 1, "is": 1, "only": 1, "test": 2, "this": 1 }`
	if strOut != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", strOut, expected)
	}
}
