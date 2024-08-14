package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("sua-chave-secreta"))

// Função para obter a sessão atual
func GetSession(r *http.Request) (*sessions.Session, error) {
	session, err := store.Get(r, "minha-sessao")
	return session, err
}

// Função para salvar a sessão
func SaveSession(w http.ResponseWriter, r *http.Request, session *sessions.Session) error {
	err := session.Save(r, w)
	return err
}

// Função para verificar se o usuário está autenticado
func IsAuthenticated(r *http.Request) bool {
	session, _ := GetSession(r)
	if session.Values["authenticated"] == nil {
		return false
	}
	return session.Values["authenticated"].(bool)
}
