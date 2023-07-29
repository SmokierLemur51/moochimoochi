package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	// "github.com/go-chi/chi/v5/middleware/cookieauth"
	// "github.com/mattn/go-sqlite3"
	// "golang.org/x/crypto/bcrypt"
)

type PageData struct {
	Title	string
	Message string
	// Account	Account
}

type Account struct {
	Username	string
	Email		string
	Password	string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "MoochiMoochi",
		Message: "The million dollar project ... ",
		// Account: account,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func processFormHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
	fmt.Printf("%s %s %s", username, email, password) // why does this not show?
}



func main() {
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.Get("/", indexHandler)
	r.Post("/process-form", processFormHandler)

	log.Println("Server started on :5000")
	http.ListenAndServe(":5000", r)
}