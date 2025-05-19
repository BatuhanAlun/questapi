package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"questapi/internal/models"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type RegisterHandlerSturct struct {
	DB *sql.DB
}

func NewRegisterHandler(db *sql.DB) *RegisterHandlerSturct {
	return &RegisterHandlerSturct{DB: db}
}

func (h *RegisterHandlerSturct) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	user.ID = uuid.New()
	user.Role = "user"

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

	_, err = h.DB.Exec("INSERT INTO users (id, username, password, role) VALUES ($1, $2, $3, $4)", user.ID, user.Username, user.Password, user.Role)

	if err != nil {
		http.Error(w, "Insert is failed", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Error on encoding new book to JSON: %v", err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
	log.Println("Succesfully served /register Post")
}
