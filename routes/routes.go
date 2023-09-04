package routes

import (

	"github.com/go-chi/chi/v5"
	"github.com/SmokierLemur51/moochimoochi/handlers"
)


func ConfigureRoutes(router *chi.Mux) {
	router.Get("/", handlers.IndexHandler)
	router.Get("/about-us", handlers.AboutHandler)
	router.Get("/contact-us", handlers.ContractingHandler)
}

