package handlers

import (
	"html/template"
	"net/http"

	"yasmim.psc.br/src/session"
)

// Função para validar o usuário (substitua pela sua lógica de validação)
func validateUser(username, password string) bool {
	// Aqui você implementaria a lógica para verificar as credenciais
	// Exemplo: consultando um banco de dados
	// ...
	return username == "tiago" && password == "senha" // Retorna true se as credenciais forem válidas
}

type loginData struct {
	Message   string
	CSRFToken string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	data := loginData{}

	if r.Method == http.MethodPost {

		submittedToken := r.FormValue("csrf_token")
		if !session.ValidateCSRFToken(w, r, submittedToken) {
			http.Error(w, "Invalid CSRF Token!", http.StatusForbidden)
			return
		}

		username := template.HTMLEscapeString(r.FormValue("username"))
		password := template.HTMLEscapeString(r.FormValue("password"))

		if username == "" || password == "" {
			data.Message = "Por favor, preencha todos os campos"
		} else if !validateUser(username, password) {
			data.Message = "Usuário ou senha inválidos"
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

	data.CSRFToken = session.GenerateCSRFToken(w, r)

	tmpl, err := template.ParseFiles("templates/base.html", "templates/_login.html")
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
