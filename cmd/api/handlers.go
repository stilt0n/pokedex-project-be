package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	// Fprint prints to a stream
	fmt.Fprint(w, "Hello, world!")
}
