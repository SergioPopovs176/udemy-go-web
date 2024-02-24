package main

import (
	"fmt"
	"net/http"
	"sergio-web-app/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	// run  go run cmd/web/*.go
	//      go run ./cmd/web
	// ctrl+C to stop programm

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/links", handlers.Links)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
