package main

import (
	"net/http"

	"github.com/go-chi/chi"
)


func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("What the fuck is up??"))
	})

	http.ListenAndServe(":5000", r)
}