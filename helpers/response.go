package helpers

import (
	"encoding/json"
	"net/http"
)

func WriteToJSON(w http.ResponseWriter, statusCode int, payload any) error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(payload)
}

