package main

import (
	"net/http"
	"html/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func main() {
	// router
	r := chi.NewRouter()
	
	// init templates
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// middleware
	r.Use(middleware.Logger)
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":5000", r)
}