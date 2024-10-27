package handlers

import (
	"github.com/domdoa/goapi/internal/middleware"
	"github.com/domdoa/goapi/internal/tools"
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func Handler(r *chi.Mux, db *tools.PostgresDb) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/users", func(router chi.Router) {
		// Middleware for /users
		router.Use(middleware.Authorization)

		router.Get("/{id}", handleGetUserByID(db))
	})
}
