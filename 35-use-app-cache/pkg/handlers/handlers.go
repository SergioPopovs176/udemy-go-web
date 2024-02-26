package handlers

import (
	"net/http"
	"sergio-web-app/pkg/config"
	"sergio-web-app/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

func (m *Repository) Links(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "links.page.tmpl")
}
