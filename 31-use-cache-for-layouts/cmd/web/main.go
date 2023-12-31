package main

// Run server
// go run cmd/web/main.go
// got run ./cmd/web

import (
	"fmt"
	"github/tsawler/go-course/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting aplication on port : %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
