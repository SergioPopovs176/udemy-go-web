package main

import (
	"github/tsawler/go-course/pkg/config"
	"github/tsawler/go-course/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
