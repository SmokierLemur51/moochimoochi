package handlers

import (	
	"net/http"
	// "encoding/json"

	"github.com/SmokierLemur51/moochimoochi/models"
	"github.com/SmokierLemur51/moochimoochi/utils"
)

const (
	TITLE string = "GutterDropper 5000"

    GUTTER_COST_FT float32 = 1.75 
    GUTTER_MARGIN float32 = 23.00
    LBS_FEET_6INCH float32 = 2.24

)

func GUTTER_OPTIONS() []models.GutterCoil { 
	return []models.GutterCoil{
	    {ID: 1, Size: "6 inches", Color: "White", OnHand: 654.45, PoundsFeetMulti: LBS_FEET_6INCH, DisplayInformation: "6\" White Gutter Coil", Cost: GUTTER_COST_FT, Margin: GUTTER_MARGIN, Selling: models.CostToSelling(GUTTER_COST_FT, GUTTER_MARGIN)},
	    {ID: 2, Size: "6 inches", Color: "Black", OnHand: 45.55, PoundsFeetMulti: LBS_FEET_6INCH, DisplayInformation: "6\" Black Gutter Coil", Cost: GUTTER_COST_FT, Margin: GUTTER_MARGIN, Selling: models.CostToSelling(GUTTER_COST_FT, GUTTER_MARGIN), },
	    {ID: 3, Size: "6 inches", Color: "Wicker", OnHand: 578.50, PoundsFeetMulti: LBS_FEET_6INCH, DisplayInformation: "6\" Wicker Gutter Coil", Cost: GUTTER_COST_FT, Margin: GUTTER_MARGIN, Selling: models.CostToSelling(GUTTER_COST_FT, GUTTER_MARGIN), },
	    {ID: 4, Size: "6 inches", Color: "Pebblestone Clay", OnHand: 120.89, PoundsFeetMulti: LBS_FEET_6INCH, DisplayInformation: "6\" Pebblestone Clay Gutter Coil", Cost: GUTTER_COST_FT, Margin: GUTTER_MARGIN, Selling: models.CostToSelling(GUTTER_COST_FT, GUTTER_MARGIN), },
	}
}



func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := models.PageData{
		Page: "index.html",
		Title: TITLE,
		CoilSelection: GUTTER_OPTIONS(),
	}

	err := utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func AboutHandler(w http.ResponseWriter, r *http.Request) {
	page := models.PageData{
		Page: "about.html",
		Title: TITLE,
		CoilSelection: GUTTER_OPTIONS(),
	}

	err := utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func ContactHandler(w http.ResponseWriter, r *http.Request) {
	page := models.PageData{
		Page: "contact.html",
		Title: TITLE,
		CoilSelection: GUTTER_OPTIONS(),
	}

	err := utils.RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
