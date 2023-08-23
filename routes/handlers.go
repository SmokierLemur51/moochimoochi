package routes

import (	
	"net/http"
	// "encoding/json"
	// "fmt"

	"github.com/SmokierLemur51/moochimoochi/models"
	"github.com/SmokierLemur51/moochimoochi/utils"
)



var TITLE string = "moochimoochi"



func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := models.PageRenderData{
		Page: "index.html",
		Title: TITLE,
	}

	err := utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
