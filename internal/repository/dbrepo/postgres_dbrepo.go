package dbrepo

import (
	"context"
	"pokedex-backend/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
)

type PostgresDBRepo struct {
	Db *pgx.Conn
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) AllPokemon() ([]*models.Pokemon, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var allPokemon []*models.Pokemon

	query := `
		select
			pokedex_id, name, type_1, coalesce(type_2, ''),
			height, weight, sprite_url
		from
			api.pokemon
		order by
			pokedex_id
	`
	rows, err := m.Db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pokemon models.PokemonRow
		err := rows.Scan(
			&pokemon.PokedexId,
			&pokemon.Name,
			&pokemon.Type1,
			&pokemon.Type2,
			&pokemon.Height,
			&pokemon.Weight,
			&pokemon.SpriteUrl,
		)
		if err != nil {
			return nil, err
		}
		converted := pokemon.ToPokemonJson()
		allPokemon = append(allPokemon, &converted)
	}

	return allPokemon, nil
}

func (m *PostgresDBRepo) Connection() *pgx.Conn {
	return m.Db
}
