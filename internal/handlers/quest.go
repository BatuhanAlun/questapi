package handlers

import (
	"encoding/json"
	"net/http"
)

func Quest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"true": "true"})
}
