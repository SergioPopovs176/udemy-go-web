package handlers

import (
	"net/http"
	"sergio-web-app/pkg/config"
	"sergio-web-app/pkg/models"
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
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// send the data to the template

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Links(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "links.page.tmpl", &models.TemplateData{})
}
