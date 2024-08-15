package handlers

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	data := templateData{}

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	executeTemplate(w, "_index", data)
}
