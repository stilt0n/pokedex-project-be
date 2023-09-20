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

// There's almost certainly a better way to handle the
// fact that the database stores Types as type_1 and type_2
// but I'm not super familiar with Go so this will be how
// I handle it for now.
type PokemonRow struct {
	PokedexId int
	Name      string
	Type1     string
	Type2     string
	Height    int
	Weight    int
	SpriteUrl string
}

func (p *PokemonRow) ToPokemonJson() Pokemon {
	var types []string
	types = append(types, p.Type1)
	if p.Type2 != "" {
		types = append(types, p.Type2)
	}
	return Pokemon{
		PokedexId: p.PokedexId,
		Name:      p.Name,
		Types:     types,
		Height:    p.Height,
		Weight:    p.Weight,
		SpriteUrl: p.SpriteUrl,
	}
}
