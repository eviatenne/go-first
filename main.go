package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Database connection string
const (
	host     = "34.75.71.40"
	port     = 5432
	user     = "postgres"
	password = "0000"
	dbname   = "pg_baby"
)

func main() {
	// Create a database connection
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Define a route handler
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// Parse URL parameters
		queryValues := r.URL.Query()
		name := queryValues.Get("name")

		// Check if the name parameter is empty
		if name == "" {
			http.Error(w, "Name parameter is required", http.StatusBadRequest)
			return
		}

		// Query the database for user with the specified name
		row := db.QueryRow("SELECT id, name FROM cites2 WHERE name = $1", name)

		// Scan the result into variables
		var id int
		var userName string
		err := row.Scan(&id, &userName)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write response
		fmt.Fprintf(w, "User found!: ID=%d, Name=%s\n", id, userName)
	})

	// Start the web server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
