package routes

import (

	"github.com/go-chi/chi/v5"
)


func ConfigureRoutes(router *chi.Mux) {
	router.Get("/", IndexHandler)
}

// should i create a separate func for api routes?