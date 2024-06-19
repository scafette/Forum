package forum

import (
	"database/sql"
	"fmt"
	"time"
)

type comment struct {
	comment_id  string
	Title       string
	Userlike    string
	Content     string
	User_id     string
	Created_at  string
	Updated_at  string
	Deleted_at  string
	Likes       int
	Dislike     int
	Liked       bool
	Disliked    bool
	Userdislike string
	Auteur      string
	post_id     string
}
type Databasecomment struct {
	ConnectedUser user
	ProfileUser   user
	Posts         []posts
	Post          posts
	comment       []comment
	Categories    []categorie
}

func CreateComment(title string, content string, post_id string, user_id string) {
	// Se connecter à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
		return
	}
	defer db.Close()

	// Préparer les dates de création et de mise à jour
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	user := GetAccount(user_id)
	// Ajouter le commentaire à la base de données
	_, err = db.Exec("INSERT INTO commentaires (comment_id, title, userlike, content, user_id, created_at, updated_at, deleted_at, likes, dislike, liked, disliked, userdislike, auteur, post_id) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		comment_id, title, "", content, user_id, currentTime, currentTime, "", 0, 0, false, false, "", user.Name, post_id)
	if err != nil {
		fmt.Printf("error inserting comment: %v", err)
	}
}

func DeleteComment(post_id string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//supprimer le commentaire de la base de données
	_, err = db.Exec("DELETE FROM comments WHERE post_id = ?", post_id)
	if err != nil {
		fmt.Printf("error deleting comment: %v", err)
	}
}
func UpdateComment(title string, content string, post_id string, user_id string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//cherche le temps pour la création et/ou update
	_ = time.Now().Format("2024-10-12 15:04:05")

	//modifier le commentaire dans la base de données
	_, err = db.Exec("UPDATE comments SET title = ?, content = ?, updated_at = ? WHERE post_id = ?", title, content, time.Now().String(), post_id)
	if err != nil {
		fmt.Printf("error updating comment: %v", err)
	}
}
func LikeComment(post_id string, user_id string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//ajouter le like à la base de données
	_, err = db.Exec("UPDATE comments SET likes = likes + 1, liked = true WHERE post_id = ?", post_id)
	if err != nil {
		fmt.Printf("error liking comment: %v", err)
	}
}
func DislikeComment(post_id string, user_id string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//ajouter le dislike à la base de données
	_, err = db.Exec("UPDATE comments SET dislike = dislike + 1, disliked = true WHERE post_id = ?", post_id)
	if err != nil {
		fmt.Printf("error disliking comment: %v", err)
	}
}
func GetAllComments() []comment {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//récupérer tous les commentaires de la base de données
	rows, err := db.Query("SELECT * FROM comments")
	if err != nil {
		fmt.Printf("error querying comments: %v", err)
	}
	defer rows.Close()

	//créer un tableau de commentaires
	var comments []comment
	for rows.Next() {
		var c comment
		err = rows.Scan(&c.Post_id, &c.Title, &c.Userlike, &c.Content, &c.User_id, &c.Created_at, &c.Updated_at, &c.Deleted_at, &c.Likes, &c.Dislike, &c.Liked, &c.Disliked, &c.Userdislike, &c.Auteur)
		if err != nil {
			fmt.Printf("error scanning comment: %v", err)
		}
		comments = append(comments, c)
	}
	return comments
}
func GetCommentsByPostID(post_id string) []comment {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//récupérer les commentaires par post_id
	rows, err := db.Query("SELECT * FROM commentaires WHERE post_id = ?", post_id)
	if err != nil {
		fmt.Printf("error querying comments: %v", err)
	}
	defer rows.Close()

	//créer un tableau de commentaires
	var comments []comment
	for rows.Next() {
		var c comment
		err = rows.Scan(&c.Post_id, &c.Title, &c.Userlike, &c.Content, &c.User_id, &c.Created_at, &c.Updated_at, &c.Deleted_at, &c.Likes, &c.Dislike, &c.Liked, &c.Disliked, &c.Userdislike, &c.Auteur)
		if err != nil {
			fmt.Printf("error scanning comment: %v", err)
		}
		comments = append(comments, c)
	}
	return comments
}
