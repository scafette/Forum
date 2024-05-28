package main

import (
	"database/sql"
	"fmt"
	forum "forum/go"
	"net/http"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		return
	}
}

func main() {
	// co a la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	checkErr(err, "Error opening database:")
	defer db.Close()

	// Crée la table users
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        customer_id INTEGER PRIMARY KEY,
        name TEXT,
        email TEXT,
        phone TEXT,
        created_at TEXT,
        updated_at TEXT,
        deleted_at TEXT,
        status TEXT,
        password TEXT,
        role TEXT,
        Avatar TEXT
    )`)
	checkErr(err, "Error creating users table:")

	// Crée la table posts
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
        post_id INTEGER PRIMARY KEY,
        title TEXT,
        userlike TEXT,
        content TEXT,
        user_id TEXT,
        created_at TEXT,
        updated_at TEXT,
        deleted_at TEXT,
        status TEXT
    )`)
	checkErr(err, "Error creating posts table:")
	Serveur()
}

func Serveur() {
	openLink()
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./src/styles"))))
	http.HandleFunc("/accueil", forum.HomePage)
	http.HandleFunc("/login", forum.LoginPage)
	http.HandleFunc("/categories", forum.CategoriesPage)
	http.HandleFunc("/dessert", forum.DessertPage)
	http.HandleFunc("/entrer", forum.EntrerPage)
	http.HandleFunc("/plat", forum.PlatPage)

	http.ListenAndServe(":2727", nil)
}

func openLink() {
	cmd := exec.Command("cmd", "/c", "start", "http://localhost:2727/accueil")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
