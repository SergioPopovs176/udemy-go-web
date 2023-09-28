package handlers

import (
	"github/tsawler/go-course/pkg/config"
	"github/tsawler/go-course/pkg/render"
	"net/http"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handler
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
