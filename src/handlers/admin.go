package handlers

import (
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	data := templateData{}

	executeTemplate(w, "_admin", data)
}
