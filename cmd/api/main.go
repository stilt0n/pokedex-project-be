package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const port = 8080

type application struct {
	Domain string
	Prod   bool
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
	// start a web server
	log.Printf("Starting application on port %d\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
