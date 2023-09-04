package handlers


import (
	"net/http"
)
 


func OrderValidationHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data...", http.StatusBadRequest)
	}
	// create order information sturct from json

	// go func() // set go routine to grab gps coord from google api

	// check coordinates against the OperationalGrid struct

	// generate order for backend if looks good
}