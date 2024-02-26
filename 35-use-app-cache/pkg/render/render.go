package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sergio-web-app/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// sets the config for the render package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	fmt.Println("Try render template", tmpl)

	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCahce()
	}

	// get requested template from cache
	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser :", err)
	}
}

func CreateTemplateCahce() (map[string]*template.Template, error) {
	// myTemplateCache := make(map[string]*template.Template)
	// или так . Это одно и то же
	myTemplateCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	// fmt.Println(pages)
	if err != nil {
		return myTemplateCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		// println(name)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myTemplateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myTemplateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myTemplateCache, err
			}
		}

		myTemplateCache[name] = ts
	}

	return myTemplateCache, nil
}

// var templateCache = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, templateName string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have template in our cache
// 	_, inMap := templateCache[templateName]
// 	if !inMap {
// 		// need to create the template
// 		log.Println("creating template " + templateName + " and adding to cache")
// 		err = createTemplateCache(templateName)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// we have template in the cache
// 		log.Println("using cached template")
// 	}

// 	tmpl = templateCache[templateName]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error executing template:", err)
// 	}
// }

// func createTemplateCache(templateName string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", templateName),
// 		"./templates/base.layout.tmpl",
// 	}

// 	// parse the template
// 	parsedTemplate, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	// add template to cache
// 	templateCache[templateName] = parsedTemplate

// 	return nil
// }
