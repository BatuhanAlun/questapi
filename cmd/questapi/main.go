package main

import (
	"net/http"
	"questapi/internal/handlers"
)

func main() {
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/quest", handlers.Quest)
	http.ListenAndServe(":8080", nil)
}
