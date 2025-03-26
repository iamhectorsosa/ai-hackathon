package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/iamhectorsosa/ai-hackathon/internal/llm"
	"github.com/iamhectorsosa/ai-hackathon/internal/models"
)

func Ask(client *llm.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var args models.AskArgs
		err := json.NewDecoder(r.Body).Decode(&args)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorReturn{Error: "Bad request"})
			return
		}

		if args.Question == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorReturn{Error: "Question is required"})
			return
		}

		answer, err := client.GetCompletion(r.Context(), args.Question)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			json.NewEncoder(w).Encode(models.ErrorReturn{Error: "Couldn't process request"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := models.AskReturn{
			Answer: answer,
		}

		json.NewEncoder(w).Encode(response)
	}
}
