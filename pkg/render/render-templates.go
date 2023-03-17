package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/mrkouhadi/go-backend-practice/pkg/config"
	"github.com/mrkouhadi/go-backend-practice/pkg/models"
)

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

var app *config.AppConfig

// NewTemplates sets the configfor the template package
func NewTemplates(appConfig *config.AppConfig) {
	app = appConfig
}

// RenderTemplate renders the requested template
func RenderTemplate(w http.ResponseWriter, tmpl string, tmplData *models.TemplateData) {
	// Get the template cache from the AppConfig
	var tmplCache map[string]*template.Template
	if app.UseCache {
		tmplCache = app.TemplateCache
	} else {
		tmplCache, _ = CreateTemplateCache()
	}

	// get requested template from cached templates
	t, ok := tmplCache[tmpl]
	if !ok {
		log.Fatal("Could not get the template from Cached templates ! ")
	}
	buf := new(bytes.Buffer)
	tmplData = AddDefaultData(tmplData)
	_ = t.Execute(buf, tmplData)

	// render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache create cache for templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}
	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		fileName := filepath.Base(page) // filepath.Base returns the last element of the path
		templSet, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// look for any layout that exist in that directory
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			templSet, err = templSet.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[fileName] = templSet
	}
	return myCache, nil
}
