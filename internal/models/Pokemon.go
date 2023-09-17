package models

import "time"

type Pokemon struct {
	PokedexId int      `json:"pokedexId"`
	Name      string   `json:"name"`
	Types     []string `json:"types"`
	Height    int      `json:"height"`
	Weight    int      `json:"weight"`
	SpriteUrl string   `json:"spriteUrl"`
	// The `json:"-"` here is saying not to include this in JSON output
	CreatedAt  time.Time `json:"-"`
	ModifiedAt time.Time `json:"-"`
}
