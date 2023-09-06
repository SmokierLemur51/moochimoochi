package handlers


import (
	"net/http"
)
 
const (
	ERROR_MSG string = "Error parsing form" 
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


// need to validate that the address coordinates are within the 
// operational range of the company
func AddressValidation(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, ERROR_MSG, http.StatusBadRequest)
	}
}