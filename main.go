package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open a database connection
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Create a table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (customer_id INTEGER PRIMARY KEY, name TEXT, email TEXT, phone TEXT, created_at TEXT, updated_at TEXT, deleted_at TEXT, status TEXT, password TEXT, role TEXT, Avatar TEXT)")
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}
	// Create a table for forum posts
	_, err = db.Exec("CREATE TABLE posts (post_id INTEGER PRIMARY KEY, title TEXT, content TEXT, user_id TEXT, created_at TEXT, updated_at TEXT, deleted_at TEXT, status TEXT)")
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}
}
