package main

import (
	"log"
	"net/http"
)

func main() {
	staticPath := "static"
	// BEGIN static server
	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// END static server
	// BEGIN root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	// END root handler
	log.Fatal(http.ListenAndServe(":8080", nil))
}
