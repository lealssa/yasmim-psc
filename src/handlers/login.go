package handlers

import (
	"html/template"
	"net/http"
)

// Função para validar o usuário (substitua pela sua lógica de validação)
func validateUser(username, password string) bool {
	// Aqui você implementaria a lógica para verificar as credenciais
	// Exemplo: consultando um banco de dados
	// ...
	return true // Retorna true se as credenciais forem válidas
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		// Parse o form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erro ao parsear formulário", http.StatusBadRequest)
			return
		}

		// username := r.FormValue("username")
		// password := r.FormValue("password")

	}

	tmpl, err := template.ParseFiles("templates/base.html", "templates/_login.html")
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
