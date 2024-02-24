package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// ctrl+C to stop programm

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/links", Links)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
