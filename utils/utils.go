package utils

import (
	"net/http"
	"html/template"
	
	"github.com/SmokierLemur51/moochimoochi/models"
)

func RenderTemplate(w http.ResponseWriter, data models.PageRenderData) error {
	tmpl, err := template.ParseFiles("templates/" + data.Page)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}

