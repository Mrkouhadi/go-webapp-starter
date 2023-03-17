package handlers

import (
	"net/http"

	"github.com/mrkouhadi/go-backend-practice/pkg/config"
	"github.com/mrkouhadi/go-backend-practice/pkg/models"
	"github.com/mrkouhadi/go-backend-practice/pkg/render"
)

// Repository is the type of repository
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository that is USED by the handlers
var Repo *Repository

// NewRepository  creates a new repository
func NewRepository(AC *config.AppConfig) *Repository {
	return &Repository{
		App: AC,
	}
}

// NewHandlers set the repository for the handlers
func NewHandlers(R *Repository) {
	Repo = R
}

// handlers home and about
// home
func (R *Repository) Home(w http.ResponseWriter, r *http.Request) {
	greet := r.RemoteAddr
	R.App.Session.Put(r.Context(), "greet", greet)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// about
func (R *Repository) About(w http.ResponseWriter, r *http.Request) {
	// create a map
	strMap := make(map[string]string)
	remoteip := R.App.Session.GetString(r.Context(), "greet")
	strMap["greet"] = remoteip
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: strMap})
}
