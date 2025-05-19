package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"questapi/internal/models"
	"strings"
)

type ChallangeDBHandler struct {
	DB *sql.DB
}

func NewChallangeHandler(db *sql.DB) *ChallangeDBHandler {
	return &ChallangeDBHandler{DB: db}
}

func (h *ChallangeDBHandler) CreateChallangeHandler(w http.ResponseWriter, r *http.Request) {

	var newChallange models.Challange
	err := json.NewDecoder(r.Body).Decode(&newChallange)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if newChallange.Name == "" {
		http.Error(w, "Challange Name Could Not be Empty", http.StatusBadRequest)
		return
	}

	_, err = h.DB.Exec("INSERT INTO challanges (challange_name) VALUES ($1)", newChallange.Name)

	if err != nil {
		http.Error(w, "Insert Failed", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newChallange)
	if err != nil {
		http.Error(w, "Challange Encode Error", http.StatusInternalServerError)
		return
	}
	log.Println("Succesfully Served /admin/challange POST")

}

func (h *ChallangeDBHandler) DeleteChallangeHandler(w http.ResponseWriter, r *http.Request) {

	pathSegments := strings.Split(r.URL.Path, "/")

	challengeID := pathSegments[3]

	_, err := h.DB.Exec("DELETE FROM challanges WHERE id = $1", challengeID)

	if err != nil {
		http.Error(w, "Delete Failed", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"challange-id": challengeID})
	if err != nil {

		http.Error(w, "JSON Encode Error", http.StatusInternalServerError)
		return
	}

	log.Println("Challange Deleted Succesfully")

}

func (h *ChallangeDBHandler) UpdateChallangeHandler(w http.ResponseWriter, r *http.Request) {

	var newChallange models.Challange
	pathSegments := strings.Split(r.URL.Path, "/")

	challangeId := pathSegments[3]

	err := json.NewDecoder(r.Body).Decode(&newChallange)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	sqlresult, err := h.DB.Exec("UPDATE challanges SET challange_name = $1 WHERE id = $2", newChallange.Name, challangeId)

	if err != nil {
		rowsAffected, sqlErr := sqlresult.RowsAffected()
		if sqlErr != nil {
			http.Error(w, "Error getting Rows", http.StatusInternalServerError)
			return
		} else if rowsAffected == 0 {
			http.Error(w, "Id Not Found", http.StatusBadRequest)
			return
		} else {
			http.Error(w, "Update Failed", http.StatusInternalServerError)
			return
		}

	}

	w.WriteHeader(http.StatusOK)
	log.Println("/admin/challanges/ PUT Succesfully Served")

}

func (h *ChallangeDBHandler) ListChallangeHandler(w http.ResponseWriter, r *http.Request) {

	var challanges []models.Challange

	rows, err := h.DB.Query("SELECT challange_name FROM challanges")

	if err != nil {
		http.Error(w, "SQL error", http.StatusInternalServerError)
		return
	}

}
