package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type Pepe struct {
	URL string `json:"url"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE_FILE := os.Getenv("DATABASE_FILE")
	PORT := os.Getenv("PORT")

	db, err := sql.Open("sqlite3", DATABASE_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Please go to api endpoint at /random")
	})

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
		}

		var pepe Pepe
		row := db.QueryRow("SELECT url FROM pepe ORDER BY RANDOM() LIMIT 1")
		err := row.Scan(&pepe.URL)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "No pepe found", http.StatusNotFound)
			} else {
				http.Error(w, "Database error", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pepe)
	})

	log.Printf("Server starting on port %s...", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
