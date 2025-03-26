package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/iamhectorsosa/ai-hackathon/internal/config"
	"github.com/iamhectorsosa/ai-hackathon/internal/handlers"
)

func New() *chi.Mux {
	return chi.NewRouter()
}

func RegisterRoutes(r *chi.Mux, cfg *config.Config) {
	r.Get("/status", handlers.Status)
	r.Get("/error", handlers.Error)

	r.Get("/posts", handlers.GetPosts)

	r.Post("/greet", handlers.Greet)
}
