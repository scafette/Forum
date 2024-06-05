package forum

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type user struct {
	Customer_id string
	Name        string
	Created_at  string
	Updated_at  string
	Deleted_at  string
	Status      string
	Password    string
	Role        string
	Avatar      string
}

var ConnectedUser user

func Signup(name string, password string, role string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//générer un id unique pour l'user
	id := uuid.New().String()

	//ajouter l'user à la base de données
	_, err = db.Exec("INSERT INTO users (customer_id, name, created_at, updated_at, deleted_at, status, password, role, avatar) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", id, name, time.Now().String(), time.Now().String(), "", "active", password, role, "")
	if err != nil {
		fmt.Printf("error inserting user: %v", err)
	}
}
func Login(name string, password string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//récupérer l'user de la base de données
	rows, err := db.Query("SELECT * FROM users WHERE name = ? AND password = ?", name, password)
	if err != nil {
		fmt.Printf("error selecting user: %v", err)
	}
	defer rows.Close()

	//stocker l'user connecté
	for rows.Next() {
		err = rows.Scan(&ConnectedUser.Customer_id, &ConnectedUser.Name, &ConnectedUser.Created_at, &ConnectedUser.Updated_at, &ConnectedUser.Deleted_at, &ConnectedUser.Status, &ConnectedUser.Password, &ConnectedUser.Role, &ConnectedUser.Avatar)
		if err != nil {
			fmt.Printf("error scanning user: %v", err)
		}
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
func Updateaccount(userID string, name string, avatar string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//met à jour les infos de l'user
	_, err = db.Exec("UPDATE users SET name = ?, email = ?, phone = ?, avatar = ? WHERE customer_id = ?", name, avatar, userID)
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
