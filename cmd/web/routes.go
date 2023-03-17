package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrkouhadi/go-backend-practice/pkg/config"
	"github.com/mrkouhadi/go-backend-practice/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	//// middlewares
	// using a middleware called Recoverer by Chi, which absorbs panics and prints the stack trace
	// it gives us usefull information about the reason of the panic to help us resolve it
	mux.Use(middleware.Recoverer)
	// csrf token
	mux.Use(NoSurf)
	// load and save session on every request
	mux.Use(LoadSession)
	// my own middleware that is only logging somethig to the console. for experimenting purpose
	mux.Use(LogToConsole)

	/// handlers
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
