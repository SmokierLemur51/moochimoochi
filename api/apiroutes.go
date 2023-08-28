package api

import (

	"github.com/go-chi/chi/v5"
)

// is this even required ? 
func API_Configure_Routes(apiRouter *chi.Mux) {
	apiRouter.Get("/schedule", ScheduleHandler)
}