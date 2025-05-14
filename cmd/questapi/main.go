package main

import (
	"net/http"
	"questapi/internal/handlers"
)

func main() {
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.ListenAndServe(":8080", nil)
}
