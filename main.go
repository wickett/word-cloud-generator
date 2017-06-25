package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
	"github.com/wickett/word-cloud-generator/wordyapi"
)

// uploadHandler converts post request body to string
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/upload.tmpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		text := r.FormValue("text")
		t := wordyapi.TextToParse{"hello", text}
		fmt.Println(string(wordyapi.ParseText(t)))
	}

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/upload", uploadHandler).Methods("POST", "GET")
	//	r.HandleFunc("/wordy", WordyHandler)
	http.Handle("/", r)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8888", r))
}
