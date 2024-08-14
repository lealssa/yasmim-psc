package session

import (
	"net/http"

	"crypto/rand"
	"encoding/base64"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("sua-chave-secreta"))

// Função para obter a sessão atual
func GetSession(r *http.Request) (*sessions.Session, error) {
	session, err := store.Get(r, "session_id")
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

func GenerateCSRFToken(w http.ResponseWriter, r *http.Request) string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	csrfToken := base64.URLEncoding.EncodeToString(b)

	session, err := GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	session.Values["csrf_token"] = csrfToken
	err = SaveSession(w, r, session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return csrfToken
}

func ValidateCSRFToken(w http.ResponseWriter, r *http.Request, csrfToken string) bool {

	session, err := GetSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return csrfToken == session.Values["csrf_token"].(string)
}
