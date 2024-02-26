package main

// Run server
// go run cmd/web/main.go
// go run ./cmd/web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SergioPopovs176/udemy-go-web/pkg/config"
	"github.com/SergioPopovs176/udemy-go-web/pkg/handlers"
	"github.com/SergioPopovs176/udemy-go-web/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
