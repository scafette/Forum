package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type user struct {
	Customer_id string
	Name        string
	Email       string
	Phone       string
	Created_at  string
	Updated_at  string
	Deleted_at  string
	Status      string
	Password    string
	Role        string
	Avatar      string
}

var ConnectedUser user

func Signup(name string, email string, phone string, password string, role string, avatar string) {
	// crée le UUID 4 pour le customer_id
	u2, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID %v", u2)

	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// cherche le temps pour la création et/ou update
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Insertion d'un utilisateur
	_, err = db.Exec("INSERT INTO users (customer_id, name, email, phone, created_at, updated_at, status, password, role, Avatar) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", u2.String(), name, email, phone, currentTime, currentTime, password, role, avatar)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return
	}
}
func Login(email string, password string) {
	// co base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// recherche de l'user avk Select ( stock tout dans le rows)
	rows, err := db.Query("SELECT * FROM users WHERE email = ? AND password = ?", email, password)
	if err != nil {
		fmt.Println("Error selecting user:", err)
		return
	}
	defer rows.Close()
	// boucle pour parcourir les rows , et stocke les valeurs dans infoUser une fois l'user connecté ( ConnectedUser)
	for rows.Next() {
		var infoUser user
		err = rows.Scan(&infoUser.Customer_id, &infoUser.Name, &infoUser.Email, &infoUser.Phone, &infoUser.Created_at, &infoUser.Updated_at, &infoUser.Deleted_at, &infoUser.Status, &infoUser.Password, &infoUser.Role, &infoUser.Avatar)
		if err != nil {
			fmt.Println("Error scanning user:", err)
			return
		}
		fmt.Printf("customer_id: %s, name: %s, email: %s, phone: %s, created_at: %s, updated_at: %s, deleted_at: %s, status: %s, password: %s, role: %s, Avatar: %s\n", infoUser.Customer_id, infoUser.Name, infoUser.Email, infoUser.Phone, infoUser.Created_at, infoUser.Updated_at, infoUser.Deleted_at, infoUser.Status, infoUser.Password, infoUser.Role, infoUser.Avatar)
		ConnectedUser = infoUser
	}
}
