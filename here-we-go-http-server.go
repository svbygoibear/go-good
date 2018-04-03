package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Creating a web-server in GoLang
	// If at the default directory, make use of the handler function
	http.HandleFunc("/", handler)

	// This is for the earth directory
	http.HandleFunc("/earth", handler2)

	// Listening on port/address and handler for this (in this case nothing)
	http.ListenAndServe(":8080", nil)

}

// Receives the request and prints a response for the client in the response writer
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\n")
}