package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

type Status struct {
	Status string `json:"status"`
}

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type GreetPayload struct {
	Message string `json:"message"`
}

type GreetReponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(100, time.Minute))

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Status{
			Status: "OK",
		})
	})

	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Houston is down (intentional)"})
	})

	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			[]Post{
				{ID: 1, Title: "Introduction to Go", Body: "Go is a statically typed, compiled language designed at Google."},
				{ID: 2, Title: "RESTful APIs with Chi", Body: "Chi is a lightweight, idiomatic and composable router for building Go HTTP services."},
				{ID: 3, Title: "Middleware in Go", Body: "Middleware is a powerful concept that allows you to add functionality to your request handling pipeline."},
				{ID: 4, Title: "JSON Serialization", Body: "Go provides excellent support for encoding and decoding JSON with the encoding/json package."},
				{ID: 5, Title: "Error Handling Patterns", Body: "Proper error handling is crucial for building robust and maintainable applications."},
			},
		)
	})

	r.Post("/greet", func(w http.ResponseWriter, r *http.Request) {
		var payload GreetPayload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		if payload.Message == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "message is required"})
			return
		}

		w.Header().Set("Content-Type", "application/json")

		response := GreetReponse{
			Message: payload.Message + " is OK!",
		}

		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", r)
}
