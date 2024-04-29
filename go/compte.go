package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func Signup(name string, email string, phone string, password string, role string, avatar string) {
	// crée le UUID 4 pour le customer_id
	u2, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID %v", u2)

	// Open a database connection
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// cherche le temps pour la création et/ou update
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Insert a user
	_, err = db.Exec("INSERT INTO users (customer_id, name, email, phone, created_at, updated_at, status, password, role, Avatar) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", u2.String(), name, email, phone, currentTime, currentTime, password, role, avatar)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return
	}
}
