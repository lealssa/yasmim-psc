package main

import (
	"fmt"
	"net/http"

	"yasmim.psc.br/src/handlers"
	"yasmim.psc.br/src/middleware"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/altera-senha", handlers.ChangePasswordHandler)
	http.HandleFunc("/lembrar-senha", handlers.ForgetPasswordHandler)
	http.HandleFunc("/blog", handlers.BlogHandler)

	http.Handle("/admin", middleware.AuthRequiredMiddleware(http.HandlerFunc(handlers.AdminHandler)))

	println("Start listening on 8080 port")
	err := http.ListenAndServe(":8080", middleware.LogMiddleware(http.DefaultServeMux))
	if err != nil {
		fmt.Printf("Error starting application: %s", err.Error())
	}
}
