package repository

import (
	"pokedex-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type DatabaseRepo interface {
	AllPokemon() ([]*models.Pokemon, error)
	Connection() *pgx.Conn
}
