package main

import (
	"encoding/json"
	"fmt"
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

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Somewhat similar to `res.send`
	w.Write(out)
}

func (app *application) AllPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := app.Db.AllPokemon()
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := json.Marshal(pokemon)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
