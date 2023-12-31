package main

import (
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Pokedex backend is running",
		Version: "0.1.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := app.Db.AllPokemon()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, pokemon)
}
