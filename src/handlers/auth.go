package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"yasmim.psc.br/src/session"
)

// Função para validar o usuário (substitua pela sua lógica de validação)
func validateUser(username, password string) bool {
	// Aqui você implementaria a lógica para verificar as credenciais
	// Exemplo: consultando um banco de dados
	// ...
	return username == "tiago" && password == "senha" // Retorna true se as credenciais forem válidas
}

type authData struct {
	Message     string
	MessageType string
	Email       string
	CSRFToken   string
}

func maskEmail(email string) string {
	if len(email) < 5 {
		return email // Email muito curto para ocultar
	}

	// Encontrar o índice do @
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return email // Email inválido
	}

	// Calcular o número de caracteres a serem ocultados
	maskLength := atIndex - 3

	// Construir a string com a máscara
	return email[:3] + strings.Repeat("*", maskLength) + email[atIndex:]
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	data := authData{}

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

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {

	data := authData{}

	if r.Method == http.MethodPost {

		submittedToken := r.FormValue("csrf_token")
		if !session.ValidateCSRFToken(w, r, submittedToken) {
			http.Error(w, "Invalid CSRF Token!", http.StatusForbidden)
			return
		}

		oldPassword := template.HTMLEscapeString(r.FormValue("old-password"))
		newPassword := template.HTMLEscapeString(r.FormValue("new-password"))
		retypePassword := template.HTMLEscapeString(r.FormValue("retype-password"))

		if oldPassword == "" || newPassword == "" || retypePassword == "" {
			data.Message = "Por favor, preencha todos os campos"
		} else if newPassword != retypePassword {
			data.Message = "As senhas não são iguais"
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

	data.CSRFToken = session.GenerateCSRFToken(w, r)

	tmpl, err := template.ParseFiles("templates/base.html", "templates/_altera_senha.html")
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

func ForgetPasswordHandler(w http.ResponseWriter, r *http.Request) {

	data := authData{}
	data.MessageType = "warning"

	if r.Method == http.MethodPost {

		submittedToken := r.FormValue("csrf_token")
		if !session.ValidateCSRFToken(w, r, submittedToken) {
			http.Error(w, "Invalid CSRF Token!", http.StatusForbidden)
			return
		}

		data.Message = "Uma nova senha foi gerada e enviada para o email cadastrado"
		data.MessageType = "success"
	}

	data.CSRFToken = session.GenerateCSRFToken(w, r)
	data.Email = maskEmail("yasmim-psc@gmail.com")

	tmpl, err := template.ParseFiles("templates/base.html", "templates/_lembrar_senha.html")
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
