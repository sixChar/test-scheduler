package main

import (
	"html/template"

	"log"
	"net/http"
)

var pageTemplates = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
  
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := pageTemplates.ExecuteTemplate(w, "root", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
