package middleware

import (
	"log"
	"net/http"

	"yasmim.psc.br/src/session"
)

func LogMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func AuthRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verifica se o usuário está autenticado
		if !session.IsAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		// Se o usuário estiver autenticado, permite o acesso à próxima função
		next.ServeHTTP(w, r)
	})
}
