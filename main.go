package main

import (
	"net/http"
	"log"

	"github.com/go-chi/chi/v5"

	"github.com/SmokierLemur51/moochimoochi/routes"
)



func main() {
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	routes.ConfigureRoutes(r)

	log.Println("Starting server on :5000")
	http.ListenAndServe(":5000", r)
}