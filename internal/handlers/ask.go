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
		if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
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

		// INFO: for a structured response the client requires tools, ref: https://docs.anthropic.com/en/docs/build-with-claude/tool-use/overview
		tool, toolChoice := llm.GenerateTool[models.AskReturn]("get_ask_return", "structured response to the question asked")
		answer, err := client.GetStructuredCompletion(
			r.Context(),
			args.Question,
			tool,
			toolChoice,
		)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			json.NewEncoder(w).Encode(models.ErrorReturn{Error: "Couldn't process request"})
			return
		}

		var response models.AskReturn
		if err := json.Unmarshal(answer, &response); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			json.NewEncoder(w).Encode(models.ErrorReturn{Error: "Couldn't process response"})
			return
		}

		if response.Answer == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			json.NewEncoder(w).Encode(models.ErrorReturn{Error: "Couldn't process response"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
