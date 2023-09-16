package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokedex-backend/internal/models"
	"time"
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
	var pokemon []models.Pokemon

	ivysaur := models.Pokemon{
		ID:         2,
		Name:       "Ivysaur",
		Types:      []string{"grass", "poison"},
		Height:     10,
		Weight:     130,
		SpriteUrl:  "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/2.png",
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	pikachu := models.Pokemon{
		ID:         25,
		Name:       "Pikachu",
		Types:      []string{"electric"},
		Height:     4,
		Weight:     60,
		SpriteUrl:  "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/25.png",
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	pokemon = append(pokemon, ivysaur, pikachu)

	out, err := json.Marshal(pokemon)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
