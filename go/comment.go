package forum

import (
	"database/sql"
	"fmt"
	"time"
)

type comment struct {
	Post_id     string
	Title       string
	Userlike    string
	Content     string
	User_id     string
	Created_at  string
	Updated_at  string
	Deleted_at  string
	Status      string
	Likes       int
	Dislike     int
	Liked       bool
	Disliked    bool
	Userdislike string
	Auteur      string
}

func CreateComment(title string, content string, post_id string, user_id string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	// cherche le temps pour la création et/ou update
	_ = time.Now().Format("2024-10-12 15:04:05")

	//ajouter le commentaire à la base de données
	_, err = db.Exec("INSERT INTO comments (post_id, title, userlike, content, user_id, created_at, updated_at, deleted_at, status, likes, dislike, liked, disliked, userdislike, auteur) VALUES (?; ?; ?; ?; ?; ?; ?; ?; ?; ?; ?; ?; ?; ?; ?)", post_id, title, "", content, user_id, time.Now().String(), time.Now().String(), "", "active", 0, 0, false, false, "", "")
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
