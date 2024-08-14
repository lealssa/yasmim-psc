package handlers

import (
	"html/template"
	"net/http"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/base.html", "templates/_blog.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
