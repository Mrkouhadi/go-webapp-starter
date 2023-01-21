package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mrkouhadi/go-simplewebapp/pkg/config"
	"github.com/mrkouhadi/go-simplewebapp/pkg/handlers"
	"github.com/mrkouhadi/go-simplewebapp/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// this should be changed to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // stays there even the session is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //which is false, i should get changed during production though

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(" Cannot Create Template Cache ")
	}
	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}
	fmt.Println("Listening to Port:8080")
	err = srv.ListenAndServe()

	log.Fatal(err)
}
