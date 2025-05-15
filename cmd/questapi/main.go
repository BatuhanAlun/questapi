package main

import (
	"net/http"
	"questapi/internal/handlers"
	"questapi/internal/utils"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("POST /register", handlers.RegisterHandler)
	router.HandleFunc("POST /login", handlers.LoginHandler)
	router.Handle("GET /quest", utils.MiddlewareAuthJWT(http.HandlerFunc(handlers.Quest)))

	http.ListenAndServe(":8080", router)
}
