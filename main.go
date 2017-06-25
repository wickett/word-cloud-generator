package main

import (
	"fmt"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
	"github.com/wickett/word-cloud-generator/wordyapi"
)

// PostHandler converts post request body to string
func PostHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	t := wordyapi.TextToParse{"hello", text}
	fmt.Println(wordyapi.ParseText(t))
}

func main() {

	x := wordyapi.TextToParse{"hello", "how are you"}

	fmt.Println(wordyapi.ParseText(x))
	r := mux.NewRouter()
	r.HandleFunc("/upload", PostHandler).Methods("POST")
	//	r.HandleFunc("/wordy", WordyHandler)
	http.Handle("/", r)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8888", r))
}
