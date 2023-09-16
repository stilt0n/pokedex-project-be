package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// set application config
	var app application
	// read from command line for flags

	// connect to the database
	app.Domain = "example.com" // this is just to silence compiler's unused variables error
	// start a web server
	log.Printf("Starting application on port %d\n", port)

	// handles a route using a function
	// this is a lot like express app.get('/', () => {})
	http.HandleFunc("/", Hello)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}

}
