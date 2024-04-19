package main

//
//import (
//	"database/sql"
//	"fmt"
//	"html/template"
//	"log"
//	"net/http"
//
//	_ "github.com/lib/pq"
//)
//
//// Database connection string
////
////	connStr := "host=34.75.71.40 port=5432 user=postgres password=0000 dbname=pg_baby sslmode=disable"
//const (
//	host     = "34.75.71.40"
//	port     = 5432
//	user     = "postgres"
//	password = "0000"
//	dbname   = "pg_baby"
//)
//
//// User struct representing a user
//type User struct {
//	ID   int
//	Name string
//}
//
//func main() {
//	// Create a database connection
//	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	// Define a route handler
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		// Query the database for users
//		rows, err := db.Query("SELECT id, name FROM cites2")
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		defer rows.Close()
//
//		// Create a slice to hold user data
//		var users []User
//
//		// Iterate over the result set
//		for rows.Next() {
//			var user User
//			if err := rows.Scan(&user.ID, &user.Name); err != nil {
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//			users = append(users, user)
//		}
//		if err := rows.Err(); err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		// Render the HTML template with user data
//		tmpl := template.Must(template.New("index").Parse(`
//			<!DOCTYPE html>
//			<html>
//			<head>
//				<title>Cities</title>
//			</head>
//			<body>
//				<h1>Users</h1>
//				<ul>
//					{{range .}}
//					<li>{{.ID}} - {{.Name}}</li>
//					{{end}}
//				</ul>
//			</body>
//			</html>
//		`))
//		if err := tmpl.Execute(w, users); err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//		}
//	})
//
//	// Start the web server
//	log.Println("Server started on :8080")
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
