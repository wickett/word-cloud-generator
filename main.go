package main

import (
	"html/template"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
	"github.com/wickett/word-cloud-generator/wordyapi"
)

// uploadHandler converts post request body to string
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	text := r.FormValue("text")
	t := wordyapi.TextToParse{Title: "hello", Text: text}
	w.Write(wordyapi.ParseText(t))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/upload.tmpl")
	t.Execute(w, nil)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api", uploadHandler).Methods("POST")
	r.HandleFunc("/", mainHandler).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8888", r))
}
