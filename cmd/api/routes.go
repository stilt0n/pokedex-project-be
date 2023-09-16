package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// Middleware
	mux.Use(middleware.Recoverer)

	// Routes
	mux.Get("/", app.Home)

	return mux
}
