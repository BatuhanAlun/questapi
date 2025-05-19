package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"questapi/internal/models"
	"questapi/internal/utils"
)

type AuthHandler struct {
	DB *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	var sqlUsername string
	var sqlPassword string
	var sqlRole string

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username or Password could not be blank!", http.StatusBadRequest)
	}

	err = h.DB.QueryRow("SELECT username, password, role FROM users WHERE username = $1", user.Username).Scan(&sqlUsername, &sqlPassword, &sqlRole)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not Found", http.StatusNotFound)
			return
		} else {
			http.Error(w, "SQL Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	if user.Username != sqlUsername || user.Password != sqlPassword {
		http.Error(w, "Credentials Not Correct", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Username, sqlRole)

	if err != nil {
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
	log.Println("Login succesful, JWT sent")
}
