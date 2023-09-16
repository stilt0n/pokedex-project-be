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
	if app.Prod {
		mux.Use(app.enableCORS)
	} else {
		mux.Use(app.enableLocalCORS)
	}
	// Routes
	mux.Get("/", app.Home)

	mux.Get("/pokemon", app.AllPokemon)

	return mux
}
