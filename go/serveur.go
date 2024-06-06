package forum

import (
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// VARIABLE DES TEMPLATES (PAGE HTML)
var Home = template.Must(template.ParseFiles("./src/templates/home.html"))
var HomeConnected = template.Must(template.ParseFiles("./src/templates/home.html"))
var Connection = template.Must(template.ParseFiles("./src/templates/login.html"))
var Categories = template.Must(template.ParseFiles("./src/templates/categories.html"))
var Dessert = template.Must(template.ParseFiles("./src/templates/dessert.html"))
var Plat = template.Must(template.ParseFiles("./src/templates/plat.html"))
var Entrer = template.Must(template.ParseFiles("./src/templates/entrer.html"))
var profile = template.Must(template.ParseFiles("./src/templates/profile.html"))
var postcreate = template.Must(template.ParseFiles("./src/templates/postcreate.html"))
var Register = template.Must(template.ParseFiles("./src/templates/register.html"))
var ErreurRegister = template.Must(template.ParseFiles("./src/templates/erreurregister.html"))

// FONCTIONS DES PAGES
func HomePage(w http.ResponseWriter, r *http.Request) {

	if ConnectedUser.Customer_id != "" {
		p := "Home page pour utilisateur connecté"
		err := HomeConnected.ExecuteTemplate(w, "home.html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return

	}

	// SI LUTILISATEUR NEST PAS CONNECTÉ
	p := "Home page pour utilisateur non connecté"
	err := Home.ExecuteTemplate(w, "home.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username") // RECUPERE LA DONNEE DE LA PAGE HTML (INPUT DE L'USER) (ID !!!!!!) ET LA STOCKE DANSnammmmmmmmmmmeeeeeee
		mdp := r.FormValue("password")      // RECUPERE LA DONNEE DE LA PAGE HTML (INPUT DE L'USER) (ID !!!!!!) ET LA STOCKE DANS MDP
		mdpConfirm := r.FormValue("confirm-password")

		if mdp == mdpConfirm {
			Signup(username, mdp, "user")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			err := ErreurRegister.ExecuteTemplate(w, "erreurregister.html", "")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	p := ""
	err := Register.ExecuteTemplate(w, "register.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username") // RECUPERE LA DONNEE DE LA PAGE HTML (INPUT DE L'USER) (ID !!!!!!) ET LA STOCKE DANS MAIL
		mdp := r.FormValue("mdp")           // RECUPERE LA DONNEE DE LA PAGE HTML (INPUT DE L'USER) (ID !!!!!!) ET LA STOCKE DANS MDP

		Login(username, mdp) // APPEL DE LA FONCTION LOGIN (voir compte.go)
	}

	p := ""
	err := Connection.ExecuteTemplate(w, "login.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	p := ""
	err := Home.ExecuteTemplate(w, "logout.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CategoriesPage(w http.ResponseWriter, r *http.Request) {

	p := ""
	err := Categories.ExecuteTemplate(w, "categories.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DessertPage(w http.ResponseWriter, r *http.Request) {
	var datas Database
	datas.Posts = GetAllPosts()
	fmt.Println(datas.Posts)
	err := Dessert.ExecuteTemplate(w, "dessert.html", datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PlatPage(w http.ResponseWriter, r *http.Request) {

	p := ""
	err := Plat.ExecuteTemplate(w, "plat.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func EntrerPage(w http.ResponseWriter, r *http.Request) {

	p := ""
	err := Entrer.ExecuteTemplate(w, "entrer.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {

	p := ""
	err := profile.ExecuteTemplate(w, "profile.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Fonction pour sauvegarder une image ( relié à la fonction PostCreatePage)
func saveImage(file multipart.File, header *multipart.FileHeader) (string, error) {
	// Crée un fichier dans le répertoire souhaité (ici "./images/")
	dst, err := os.Create("./images/" + header.Filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copie le contenu du fichier téléchargé dans le nouveau fichier
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return "./images/" + header.Filename, nil
}

func PostCreatePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")     // RECUPERE LA DONNEE DE LA PAGE HTML (INPUT DE L'USER) (ID !!!!!!)
		content := r.FormValue("content") // RECUPERE LA DONNEE DE LA PAGE HTML (INPUT DE L'USER) (ID !!!!!!)
		categories := r.FormValue("Categories")
		sub := r.FormValue("sub")
		user_id := ConnectedUser.Customer_id
		// Récupère le fichier image du formulaire
		file, handler, err := r.FormFile("image")
		if err != nil {
			fmt.Fprintf(w, "Error retrieving file: %v", err)
			return
		}
		defer file.Close()

		// Sauvegarde l'image téléchargée
		image, err := saveImage(file, handler)
		if err != nil {
			fmt.Fprintf(w, "Error saving image: %v", err)
			return
		}
		CreatePost(title, content, user_id, categories, sub, image) // APPEL DE LA FONCTION CREATEPOST (voir post.go)
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	} else {
		p := ""
		err := postcreate.ExecuteTemplate(w, "postcreate.html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	}

}
