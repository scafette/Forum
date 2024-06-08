package forum

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type posts struct {
	Post_id    string
	Title      string
	Userlike   string
	Content    string
	User_id    string
	Created_at string
	Updated_at string
	Deleted_at string
	Status     string
	Categories string
	Sub        string
	Image      string
	Likes      int
	Dislike    int
}

type Database struct {
	ConnectedUser user
	Posts         []posts
	Post          posts
	// comment		 []comments

}

func CreatePost(title string, content string, user_id string, categories string, sub string, image string) {
	// crée le UUID 4 pour le post_id
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
	currentTime := time.Now().Format("2024-10-12 15:04:05")

	// Insertion d'un post
	_, err = db.Exec("INSERT INTO posts (post_id, title, userlike, content, user_id, created_at, updated_at, deleted_at, status, categories, sub, Image, likes, dislike ) VALUES ( ?, ?, ?, ?, ?, ? ,?, ?, ?, ?, ?, ?, ?, ?)",
		u2.String(), title, "", content, user_id, currentTime, currentTime, "", "published", categories, sub, image, 0, 0)
	if err != nil {
		fmt.Println("Error inserting post:", err)
		return
	}
}
func UpdatePost(post_id string, title string, content string) {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// cherche le temps pour la création et/ou update
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Mise à jour d'un post
	_, err = db.Exec("UPDATE posts SET title = ?, content = ?, updated_at = ? WHERE post_id = ?", title, content, currentTime, post_id)
	if err != nil {
		fmt.Println("Error updating post:", err)
		return
	}
}
func DeletePost(post_id string) {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// cherche le temps pour la création et/ou update
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Suppression d'un post
	_, err = db.Exec("UPDATE posts SET deleted_at = ?, status = ? WHERE post_id = ?", currentTime, "deleted", post_id)
	if err != nil {
		fmt.Println("Error deleting post:", err)
		return
	}
}
func GetPost(post_id string) posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return posts{}
	}
	defer db.Close()

	// Recherche d'un post
	row := db.QueryRow("SELECT * FROM posts WHERE post_id = ?", post_id)
	var post posts
	err = row.Scan(&post.Post_id, &post.Title, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status)
	if err != nil {
		fmt.Println("Error getting post:", err)
		return posts{}
	}
	return post
}
func GetAllPosts() []posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return []posts{}
	}
	defer db.Close()

	// Recherche de tous les posts
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Println("Error getting posts:", err)
		return []posts{}
	}
	defer rows.Close()

	// Crée un tableau de posts
	var allPosts []posts
	for rows.Next() {
		var post posts
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike)
		if err != nil {
			fmt.Println("Error scanning post:", err)
			return []posts{}
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}

func GetPostsByUser(user_id string) []posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return []posts{}
	}
	defer db.Close()

	// Recherche de tous les posts d'un utilisateur
	rows, err := db.Query("SELECT * FROM posts WHERE user_id = ?", user_id)
	if err != nil {
		fmt.Println("Error getting posts:", err)
		return []posts{}
	}
	defer rows.Close()

	// Crée un tableau de posts
	var allPosts []posts
	for rows.Next() {
		var post posts
		err = rows.Scan(&post.Post_id, &post.Title, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status)
		if err != nil {
			fmt.Println("Error scanning post:", err)
			return []posts{}
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}
func getAllPostsDessert() []posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return []posts{}
	}
	defer db.Close()

	// Recherche de tous les posts
	rows, err := db.Query("SELECT * FROM posts WHERE categories = ?", "Dessert")
	if err != nil {
		fmt.Println("Error getting posts:", err)
		return []posts{}
	}
	defer rows.Close()

	// Crée un tableau de posts
	var allPosts []posts
	for rows.Next() {
		var post posts
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike)
		if err != nil {
			fmt.Println("Error scanning post:", err)
			return []posts{}
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}
func getAllPostsPlat() []posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return []posts{}
	}
	defer db.Close()

	// Recherche de tous les posts
	rows, err := db.Query("SELECT * FROM posts WHERE categories = ?", "Plat")
	if err != nil {
		fmt.Println("Error getting posts:", err)
		return []posts{}
	}
	defer rows.Close()

	// Crée un tableau de posts
	var allPosts []posts
	for rows.Next() {
		var post posts
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike)
		if err != nil {
			fmt.Println("Error scanning post:", err)
			return []posts{}
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}
func getAllPostsEntrer() []posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return []posts{}
	}
	defer db.Close()

	// Recherche de tous les posts
	rows, err := db.Query("SELECT * FROM posts WHERE categories = ?", "Entree")
	if err != nil {
		fmt.Println("Error getting posts:", err)
		return []posts{}
	}
	defer rows.Close()

	// Crée un tableau de posts
	var allPosts []posts
	for rows.Next() {
		var post posts
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike)
		if err != nil {
			fmt.Println("Error scanning post:", err)
			return []posts{}
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}
func LikePosts(post_id string, user_id string) {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Insertion d'un like
	_, err = db.Exec("INSERT INTO likes (post_id, user_id) VALUES (?, ?)", post_id, user_id)
	if err != nil {
		fmt.Println("Error liking post:", err)
		return
	}
}
func DislikePosts(post_id string, user_id string) {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Suppression d'un like
	_, err = db.Exec("DELETE FROM likes WHERE post_id = ? AND user_id = ?", post_id, user_id)
	if err != nil {
		fmt.Println("Error disliking post:", err)
		return
	}
}
func getlikebyUser(user_id string) bool {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return false
	}
	defer db.Close()

	// Recherche d'un like
	row := db.QueryRow("SELECT * FROM posts WHERE userlike LIKE ?", "%"+user_id+"%")
	var like struct{}
	err = row.Scan(&like)
	if err != nil {
		fmt.Println("Error getting like:", err)
		return false
	}
	return true
}
func getDislikebyUser(user_id string) bool {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return false
	}
	defer db.Close()

	// Recherche d'un dislike
	row := db.QueryRow("SELECT * FROM posts WHERE userlike NOT LIKE ?", "%"+user_id+"%")
	var dislike struct{}
	err = row.Scan(&dislike)
	if err != nil {
		fmt.Println("Error getting dislike:", err)
		return false
	}
	return true
}
