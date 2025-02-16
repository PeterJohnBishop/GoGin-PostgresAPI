package routes

import (
	"encoding/json"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {

	response := map[string]interface{}{
		"message": "Welcome to my Go/HTTP server!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
