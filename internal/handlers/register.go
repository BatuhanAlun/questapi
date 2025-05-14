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

var db *sql.DB

func InitDB() {
	connStr := "user=questapi dbname=questapi password=securepass host=localhost sslmode=disable"

	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error on opening DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error on connectiong DB: %v", err)
	}
	log.Println("Connection Succesfuly Established")

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	InitDB()
	defer db.Close()

	var user models.User

	user.ID = uuid.New()

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

	_, err = db.Exec("INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", user.ID, user.Username, user.Password)

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
