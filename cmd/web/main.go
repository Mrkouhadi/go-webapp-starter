package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mrkouhadi/go-backend-practice/pkg/config"
	"github.com/mrkouhadi/go-backend-practice/pkg/handlers"
	"github.com/mrkouhadi/go-backend-practice/pkg/render"
)

const portNumber = ":8080"

var AppConfiguration config.AppConfig
var session *scs.SessionManager

func main() {
	// set production mode to false
	AppConfiguration.InProduction = false
	// set up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // stays there even the session is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = AppConfiguration.InProduction // should get changed during production (https vs http)
	AppConfiguration.Session = session
	// creating templates cache
	tmplCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create templates cache !")
	}
	// assign new cache
	AppConfiguration.TemplateCache = tmplCache
	// set it to not reading from cache (cuz we're still in development mode)
	AppConfiguration.UseCache = false // if it's false, changes in templates will be shown without stopping the server and run it again; which means it's not reading from the cache(development mode)
	// initializing a new repository
	repo := handlers.NewRepository(&AppConfiguration)
	handlers.NewHandlers(repo)

	// let's give render pkg an access to our appConfiguration for rendering new templates
	render.NewTemplates(&AppConfiguration)

	// create a server and run it
	ourServer := &http.Server{
		Addr:    portNumber,
		Handler: routes(&AppConfiguration),
	}
	err = ourServer.ListenAndServe()
	log.Fatal(err)
}
