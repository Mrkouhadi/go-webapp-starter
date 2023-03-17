package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// LogToConsole is just for experimenting middlewares in go
func LogToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am a experimenting middlewares (LogToConsole)....")
		next.ServeHTTP(w, r)
	})
}

// NoSurf add CSRF protection to Post requests
func NoSurf(next http.Handler) http.Handler {
	// create a csrf handler
	CSRFHandler := nosurf.New(next)
	// set te base
	CSRFHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   AppConfiguration.InProduction, // will be changed in production cuz in development we don't use HTTPS we only http
		SameSite: http.SameSiteLaxMode,
	})
	return CSRFHandler
}

// LoadSession loads and save session on every request
func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
