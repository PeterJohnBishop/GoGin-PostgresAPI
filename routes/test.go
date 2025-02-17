package routes

import (
	"encoding/json"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	response := map[string]interface{}{
		"message": "Welcome to my Go/HTTP server!",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
