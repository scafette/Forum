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
func ChangePassword(userID string, newPassword string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//mettre à jour le mot de passe de l'user
	_, err = db.Exec("UPDATE users SET password = ? WHERE customer_id = ?", newPassword, userID)
	if err != nil {
		fmt.Printf("error updating password: %v", err)
	}
}
func EmailExist(email string) bool {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//cherche de l'email dans la base de données
	rows, err := db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		fmt.Printf("error selecting user: %v", err)
	}
	defer rows.Close()

	//vérifie si l'email existe
	return rows.Next()
}
func PhoneExist(phone string) bool {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//cherche le numéro de téléphone dans la base de données
	rows, err := db.Query("SELECT * FROM users WHERE phone = ?", phone)
	if err != nil {
		fmt.Printf("error selecting user: %v", err)
	}
	defer rows.Close()

	//vérifie si le numéro de téléphone existe
	return rows.Next()
}
func Deleteaccount(userID string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//supprime le compte de l'user
	_, err = db.Exec("DELETE FROM users WHERE customer_id = ?", userID)
	if err != nil {
		fmt.Printf("error deleting user: %v", err)
	}
}
func Updateaccount(userID string, name string, email string, phone string, avatar string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//met à jour les infos de l'user
	_, err = db.Exec("UPDATE users SET name = ?, email = ?, phone = ?, avatar = ? WHERE customer_id = ?", name, email, phone, avatar, userID)
	if err != nil {
		fmt.Printf("error updating user: %v", err)
	}
}
func UpdateUserRole(userID string, role string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//met à jour le rôle de l'user
	_, err = db.Exec("UPDATE users SET role = ? WHERE customer_id = ?", role, userID)
	if err != nil {
		fmt.Printf("error updating user: %v", err)
	}
}

func GetUsers() []user {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//cherche tous les utilisateurs
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Printf("error selecting users: %v", err)
	}
	defer rows.Close()

	//stocke les utilisateurs dans un tableau
	var users []user
	for rows.Next() {
		var infoUser user
		err = rows.Scan(&infoUser.Customer_id, &infoUser.Name, &infoUser.Email, &infoUser.Phone, &infoUser.Created_at, &infoUser.Updated_at, &infoUser.Deleted_at, &infoUser.Status, &infoUser.Password, &infoUser.Role, &infoUser.Avatar)
		if err != nil {
			fmt.Printf("error scanning user: %v", err)
		}
		users = append(users, infoUser)
	}
	return users
}
