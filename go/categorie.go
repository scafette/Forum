package forum

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type categorie struct {
	categorie_id string
	name         string
}

func CreateCategorie(name string) {
	u2, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO categories (categorie_id, name) VALUES (?, ?)", u2.String(), name)
	if err != nil {
		fmt.Printf("error inserting categorie: %v", err)
	}
}
