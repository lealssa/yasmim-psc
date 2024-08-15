package handlers

import (
	"html/template"
	"net/http"
)

type templateData struct {
	Message         string
	MessageType     string
	Email           string
	CSRFToken       string
	IsAuthenticated bool
}

func executeTemplate(w http.ResponseWriter, templateName string, data templateData) {
	tmpl, err := template.ParseFiles("templates/base.html", "templates/"+templateName+".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
