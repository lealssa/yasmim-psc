package handlers

import (
	"net/http"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {

	data := templateData{}

	executeTemplate(w, "_blog", data)
}
