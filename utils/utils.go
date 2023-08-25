package utils

import (
	"net/http"
	"html/template"
	
	"github.com/SmokierLemur51/moochimoochi/models"
)

func RenderTemplate(w http.ResponseWriter, data models.PageData) error {
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


func RenderAdministrativeTemplate(w http.ResponseWriter, data models.AdministrativeFeed) error {
	tmpl, err := template.ParseFiles("templates/admin/" + data.Page)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil 
}
