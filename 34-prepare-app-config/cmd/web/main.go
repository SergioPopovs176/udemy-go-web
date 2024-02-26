package main

import (
	"fmt"
	"log"
	"net/http"
	"sergio-web-app/pkg/config"
	"sergio-web-app/pkg/handlers"
	"sergio-web-app/pkg/render"
)

const portNumber = ":8080"

func main() {
	// run  go run cmd/web/*.go
	//      go run ./cmd/web
	// ctrl+C to stop programm

	var app config.AppConfig

	tc, err := render.CreateTemplateCahce()
	if err != nil {
		log.Fatal(("cannot create template cache"))
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/links", handlers.Links)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
