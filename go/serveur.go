package forum

import (
	"fmt"
	"html/template"
	"net/http"
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

// FONCTIONS DES PAGES
func HomePage(w http.ResponseWriter, r *http.Request) {

	if ConnectedUser.Email != "" { // SI LUTILISATEUR EST CONNECTÉ
		p := "Home page pour utilisateur non connecté"
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
	p := ""
	err := Connection.ExecuteTemplate(w, "register.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {

	mail := r.FormValue("mail") // RECUPERE LA DONNEE DE LA PAGE HTML (INPUT DE L'USER) (ID !!!!!!) ET LA STOCKE DANS MAIL
	mdp := r.FormValue("mdp")   // RECUPERE LA DONNEE DE LA PAGE HTML (INPUT DE L'USER) (ID !!!!!!) ET LA STOCKE DANS MDP

	fmt.Println(mail, mdp)

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

	p := ""
	err := Dessert.ExecuteTemplate(w, "dessert.html", p)
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

func PostCreatePage(w http.ResponseWriter, r *http.Request) {

	p := ""
	err := postcreate.ExecuteTemplate(w, "postcreate.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
