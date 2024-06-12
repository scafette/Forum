package forum

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

type posts struct {
	Post_id     string
	Title       string
	Userlike    string
	Content     string
	User_id     string
	Created_at  string
	Updated_at  string
	Deleted_at  string
	Status      string
	Categories  string
	Sub         string
	Image       string
	Likes       int
	Dislike     int
	Liked       bool
	Disliked    bool
	Userdislike string
}

type Database struct {
	ConnectedUser user
	Posts         []posts
	Post          posts
	// comment		 []comments
	Categories []categorie
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
	_, err = db.Exec("INSERT INTO posts (post_id, title, userlike, content, user_id, created_at, updated_at, deleted_at, status, categories, sub, Image, likes, dislike, userdislike ) VALUES ( ?, ?, ?, ?, ?, ?, ? ,?, ?, ?, ?, ?, ?, ?, ?)",
		u2.String(), title, "", content, user_id, currentTime, currentTime, "", "published", categories, sub, image, 0, 0, "")
	if err != nil {
		fmt.Println("Error inserting post:", err)
		return
	}
}

func EditPost(post_id string, title string, content string) {
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
	// currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Suppression d'un post
	_, err = db.Exec("DELETE FROM posts WHERE post_id = ?", post_id)
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
	err = row.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
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
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
		if err != nil {
			fmt.Println("Error scanning post:", err)
			return []posts{}
		}

		// if ConnectedUser.Name != "" {
		// 	if strings.Contains(post.Userlike, ConnectedUser.Name) {
		// 		post.Liked = true
		// 	} else if strings.Contains(post.Userdislike, ConnectedUser.Name) {
		// 		post.Disliked = true
		// 	}
		// }

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
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
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
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
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
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
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
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
		if err != nil {
			fmt.Println("Error scanning post:", err)
			return []posts{}
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}
func LikePost(post_id string, user_id string) {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Vérifie si l'utilisateur a déjà aimé le post
	row := db.QueryRow("SELECT * FROM posts WHERE post_id = ? AND userlike LIKE ?", post_id, "%"+user_id+"%")
	var post posts
	err = row.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
	if err != nil {
		// L'utilisateur n'a pas encore aimé le post
		_, err = db.Exec("UPDATE posts SET userlike = ? WHERE post_id = ?", post.Userlike+user_id+",", post_id)
		if err != nil {
			log.Fatal(err)
		}
		// Query qui rajoute a post.Like +1
		_, err = db.Exec("UPDATE posts SET likes = likes + 1 WHERE post_id = ?", post_id)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		// L'utilisateur a déjà aimé le post, supprime son like
		_, err = db.Exec("UPDATE posts SET userlike = ? WHERE post_id = ?", strings.ReplaceAll(post.Userlike, user_id+",", ""), post_id)
		if err != nil {
			log.Fatal(err)
		}

		// Query qui retire a post.Like -1
		_, err = db.Exec("UPDATE posts SET likes = likes - 1 WHERE post_id = ?", post_id)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func DislikePost(post_id string, user_id string) {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Vérifie si l'utilisateur a déjà aimé le post
	row := db.QueryRow("SELECT * FROM posts WHERE post_id = ? AND userdislike LIKE ?", post_id, "%"+user_id+"%")
	var post posts
	err = row.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
	if err != nil {
		// L'utilisateur n'a pas encore aimé le post
		_, err = db.Exec("UPDATE posts SET userdislike = ? WHERE post_id = ?", post.Userdislike+user_id+",", post_id)
		if err != nil {
			log.Fatal(err)
		}
		// Query qui rajoute a post.Like +1
		_, err = db.Exec("UPDATE posts SET dislike = dislike + 1 WHERE post_id = ?", post_id)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		// L'utilisateur a déjà aimé le post, supprime son like
		_, err = db.Exec("UPDATE posts SET userdislike = ? WHERE post_id = ?", strings.ReplaceAll(post.Userdislike, user_id+",", ""), post_id)
		if err != nil {
			log.Fatal(err)
		}

		// Query qui retire a post.Like -1
		_, err = db.Exec("UPDATE posts SET dislike = dislike - 1 WHERE post_id = ?", post_id)
		if err != nil {
			log.Fatal(err)
		}
	}
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
func getPostbyID(post_id string) posts {
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
	err = row.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
	if err != nil {
		fmt.Println("Error getting post:", err)
		return posts{}
	}
	return post
}

func GetAllPostsLiked(user_id string) []posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Recherche de tous les posts
	rows, err := db.Query("SELECT * FROM posts WHERE userlike LIKE ?", "%"+user_id+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Crée un tableau de posts
	var allPosts []posts
	for rows.Next() {
		var post posts
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
		if err != nil {
			log.Fatal(err)
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}

func GetAllPostsDisliked(user_id string) []posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Recherche de tous les posts
	rows, err := db.Query("SELECT * FROM posts WHERE userdislike LIKE ?", "%"+user_id+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Crée un tableau de posts
	var allPosts []posts
	for rows.Next() {
		var post posts
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
		if err != nil {
			log.Fatal(err)
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}
func getPostByCategories(subcategorie string) []posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Recherche de tous les posts
	rows, err := db.Query("SELECT * FROM posts WHERE sub = ?", subcategorie)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Crée un tableau de posts
	var allPosts []posts
	for rows.Next() {
		var post posts
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
		if err != nil {
			log.Fatal(err)
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}

func getAllDessertPostsBySubcategories(subcategorie string) []posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Recherche de tous les posts
	rows, err := db.Query("SELECT * FROM posts WHERE sub = ? AND categories = ?", subcategorie, "Dessert")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Crée un tableau de posts
	var allPosts []posts
	for rows.Next() {
		var post posts
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
		if err != nil {
			log.Fatal(err)
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}

func getAllPlatPostsBySubcategories(subcategorie string) []posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Recherche de tous les posts
	rows, err := db.Query("SELECT * FROM posts WHERE sub = ? AND categories = ?", subcategorie, "Plat")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Crée un tableau de posts
	var allPosts []posts
	for rows.Next() {
		var post posts
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
		if err != nil {
			log.Fatal(err)
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}

func getAllEntreePostsBySubcategories(subcategorie string) []posts {
	// Ouvre une connexion à la base de données
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Recherche de tous les posts
	rows, err := db.Query("SELECT * FROM posts WHERE sub = ? AND categories = ?", subcategorie, "Entree")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Crée un tableau de posts
	var allPosts []posts
	for rows.Next() {
		var post posts
		err = rows.Scan(&post.Post_id, &post.Title, &post.Userlike, &post.Content, &post.User_id, &post.Created_at, &post.Updated_at, &post.Deleted_at, &post.Status, &post.Categories, &post.Sub, &post.Image, &post.Likes, &post.Dislike, &post.Userdislike)
		if err != nil {
			log.Fatal(err)
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}
