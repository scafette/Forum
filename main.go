package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open a database connection
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Create a table
	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY,name TEXT)")
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}
}
