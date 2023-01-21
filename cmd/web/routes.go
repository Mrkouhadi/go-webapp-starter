package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrkouhadi/go-simplewebapp/pkg/config"
	"github.com/mrkouhadi/go-simplewebapp/pkg/handlers"
)

func Routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	// use middleware (through chi)
	mux.Use(middleware.Recoverer)

	mux.Use(Nosurf)
	mux.Use(LoadSession)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/features", handlers.Repo.Features)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
