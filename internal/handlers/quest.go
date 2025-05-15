package handlers

import (
	"encoding/json"
	"net/http"
	"questapi/internal/utils"
)

func Quest(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		http.Error(w, "Authorization header is requried", http.StatusUnauthorized)
		return
	}

	err := utils.AuthJWT(tokenString)

	if err != nil {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"true": "true"})
}
