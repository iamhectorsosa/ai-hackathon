package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/iamhectorsosa/ai-hackathon/internal/models"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	var payload models.GreetArgs
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorReturn{Error: err.Error()})
		return
	}

	if payload.Message == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorReturn{Error: "message is required"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := models.GreetReturn{
		Message: payload.Message + " is OK!",
	}
	json.NewEncoder(w).Encode(response)
}
