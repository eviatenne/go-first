package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Replace the connection string with your actual PostgreSQL connection details
	connStr := "host=34.75.71.40 port=5432 user=postgres password=0000 dbname=pg_baby sslmode=disable"

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Execute a query
	rows, err := db.Query("SELECT id, name FROM cites2")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var i int = 0

		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
		i++
		if i == 100 {
			break
		}

	}

	// Check for errors from iterating over the result set
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
