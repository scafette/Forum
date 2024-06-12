package forum

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type categorie struct {
	Categorie_id string
	Name         string
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

func getallcategories() []categorie {

	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		fmt.Printf("error getting categories: %v", err)
	}
	defer rows.Close()

	var categories []categorie
	for rows.Next() {
		var c categorie
		err = rows.Scan(&c.Categorie_id, &c.Name)
		if err != nil {
			fmt.Printf("error scanning categories: %v", err)
		}
		categories = append(categories, c)
	}
	return categories
}
