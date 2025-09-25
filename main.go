package main

import (
	"html/template"
	"log"
	"net/http"
)

var pageTemplates = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// BEGIN root-template
		if err := pageTemplates.ExecuteTemplate(w, "root", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// END root-template
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
