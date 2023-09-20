package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"pokedex-backend/internal/repository"
	"pokedex-backend/internal/repository/dbrepo"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

const port = 8080

type application struct {
	Domain string
	Prod   bool
	Db     repository.DatabaseRepo
}

type pgEnv struct {
	user, pass, db string
}

func loadEnv(prod bool) (pgEnv, error) {
	if prod {
		return pgEnv{}, errors.New("production version has not been implemented")
	} else {
		err := godotenv.Load("./local-dev-env/.env-local")
		if err != nil {
			return pgEnv{}, err
		}
		return pgEnv{
			user: os.Getenv("POSTGRES_USER"),
			pass: os.Getenv("POSTGRES_PASSWORD"),
			db:   os.Getenv("POSTGRES_DB"),
		}, nil
	}
}

func main() {
	// set application config
	var app application
	// read from command line for flags
	if len(os.Args) > 1 {
		if os.Args[1] == "--dev" {
			log.Println("Running in development mode")
			app.Prod = false
		} else {
			app.Prod = true
		}
	}
	// connect to the database
	app.Domain = "example.com" // this is just to silence compiler's unused variables error
	connection, err := loadEnv(app.Prod)
	if err != nil {
		log.Fatal(err)
	}
	// This will need to be changed when we introduce production databases
	databaseUrl := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", connection.user, connection.pass, connection.db)
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	if !app.Prod {
		log.Printf("Successfully connected to %s as %s\n", connection.db, connection.user)
	}
	// start a web server
	log.Printf("Starting application on port %d\n", port)
	app.Db = &dbrepo.PostgresDBRepo{Db: conn}
	defer app.Db.Connection().Close(context.Background())

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatalln(err)
	}

}
