package main

import (
	"net/http"
	"questapi/internal/database"
	"questapi/internal/handlers"
	"questapi/internal/middleware"
	"questapi/internal/utils"
)

func main() {

	db := database.InitDB()
	defer db.Close()

	authHandler := handlers.NewAuthHandler(db)
	registerHandler := handlers.NewRegisterHandler(db)
	ChallangeHandler := handlers.NewChallangeHandler(db)

	router := http.NewServeMux()

	router.HandleFunc("POST /register", registerHandler.RegisterHandler)
	router.HandleFunc("POST /login", authHandler.LoginHandler)
	router.Handle("GET /quest", utils.MiddlewareAuthJWT(http.HandlerFunc(handlers.Quest)))
	router.Handle("POST /admin/challange", middleware.AuthAdmin(http.HandlerFunc(ChallangeHandler.CreateChallangeHandler)))
	router.Handle("DELETE /admin/challange/{id}", middleware.AuthAdmin(http.HandlerFunc(ChallangeHandler.DeleteChallangeHandler)))
	router.Handle("PUT /admin/challange/{id}", middleware.AuthAdmin(http.HandlerFunc(ChallangeHandler.UpdateChallangeHandler)))

	http.ListenAndServe(":8080", router)
}
