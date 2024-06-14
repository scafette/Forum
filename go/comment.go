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
func DeleteComment(id string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//supprimer le commentaire de la base de données
	_, err = db.Exec("UPDATE comments SET status = ? WHERE id = ?", "deleted", id)
	if err != nil {
		fmt.Printf("error deleting comment: %v", err)
	}
}
func LikeComment(id string, user_id string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//ajouter le like à la base de données
	_, err = db.Exec("UPDATE comments SET likes = likes + 1 WHERE id = ?", id)
	if err != nil {
		fmt.Printf("error liking comment: %v", err)
	}
}
func DislikeComment(id string, user_id string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//ajouter le dislike à la base de données
	_, err = db.Exec("UPDATE comments SET dislike = dislike + 1 WHERE id = ?", id)
	if err != nil {
		fmt.Printf("error disliking comment: %v", err)
	}
}
func UpdateComment(id string, title string, content string, user_id string) {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//mettre à jour le commentaire de la base de données
	_, err = db.Exec("UPDATE comments SET title = ?, content = ?, updated_at = ? WHERE id = ?", title, content, time.Now().String(), id)
	if err != nil {
		fmt.Printf("error updating comment: %v", err)
	}
}
func GetComment(id string) comment {
	//co à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	//récupérer le commentaire de la base de données
	rows, err := db.Query("SELECT * FROM comments WHERE id = ?", id)
	if err != nil {
		fmt.Printf("error getting comment: %v", err)
	}
	defer rows.Close()

	var c comment
	for rows.Next() {
		err := rows.Scan(&c.Post_id, &c.Title, &c.Userlike, &c.Content, &c.User_id, &c.Created_at, &c.Updated_at, &c.Deleted_at, &c.Status, &c.Likes, &c.Dislike, &c.Liked, &c.Disliked, &c.Userdislike, &c.Auteur)
		if err != nil {
			fmt.Printf("error scanning comment: %v", err)
		}
	}

	return c
}
