package main

import (
	"net/http"
	"sergio-web-app/pkg/config"
	"sergio-web-app/pkg/handlers"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/links", http.HandlerFunc(handlers.Repo.Links))

	return mux
}
