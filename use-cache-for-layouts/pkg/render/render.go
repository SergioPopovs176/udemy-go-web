package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache() // страно, получается каждый запрос будет генерировать этот кеш ? посмотрим что будет дальше
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		fmt.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./teplates/*.pages.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./teplates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./teplates/*.layout.tmpl")
		}
		if err != nil {
			return myCache, err
		}

		myCache[name] = ts
	}

	return myCache, nil
}
