package main

import (
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	// Fprint prints to a stream
	fmt.Fprintf(w, "Hello, world from %s!", app.Domain)
}
