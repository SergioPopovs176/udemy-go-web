package main

// Run server
// go run cmd/web/main.go
// go run ./cmd/web

import (
	"fmt"
	"github/tsawler/go-course/pkg/config"
	"github/tsawler/go-course/pkg/handlers"
	"github/tsawler/go-course/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting aplication on port : %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
