package routes

import (	
	"net/http"
	// "encoding/json"

	"github.com/SmokierLemur51/moochimoochi/models"
	"github.com/SmokierLemur51/moochimoochi/utils"
)



var TITLE string = "moochimoochi"



func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := models.PageData{
		Page: "index.html",
		Title: TITLE,
	}

	err := utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


// Post data 
func OrderValidationHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data...", http.StatusBadRequest)
	}
	// create order information sturct from json

	// go func() // set go routine to grab gps coord from google api

	// check coordinates against the OperationalGrid struct

	// generate order for backend if looks good
}