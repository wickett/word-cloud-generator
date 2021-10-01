package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wickett/word-cloud-generator/wordyapi"
)

// serves up our static content like html
//go:embed static/*
var staticFiles embed.FS

func getStaticFiles() http.FileSystem {
	fsys, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}

// TextSubmission is a json title and string to submit
type TextSubmission struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/upload.tmpl")
	t.Execute(w, nil)
}

// Test json with curl using this:
// curl -H "Content-Type: application/json" -d '{"text":"this is a really really really important thing this is"}' http://localhost:8888/newapi

func receiveJSONHandler(w http.ResponseWriter, r *http.Request) {
	var textIn TextSubmission

	// don't allow huge uploads
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	//fmt.Printf(string(body))
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &textIn); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	t := wordyapi.TextToParse{Title: textIn.Title, Text: textIn.Text}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(wordyapi.ParseText(t))
}

func main() {

	staticFs := http.FileServer(getStaticFiles())

	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/api", receiveJSONHandler)
	mux.Handle("/", staticFs)

	// Bind to a port and pass our router in
	err := http.ListenAndServe(":8888", mux)

	if err != nil {
		log.Fatal(err)
	}

}
