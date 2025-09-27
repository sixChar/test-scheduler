package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

const staticPath = "static"

var pageTemplates = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	db, err := sql.Open("sqlite", "file:app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := initUsersTable(db); err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir(staticPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := pageTemplates.ExecuteTemplate(w, "root", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initUsersTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                email TEXT NOT NULL UNIQUE,
                password_hash TEXT NOT NULL
        )`)
	return err
}
