package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	fmt.Println("Try render template", tmpl)

	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	// fmt.Println(parsedTemplate)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing template:", err)
	}
}

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have template in our cache
	_, inMap := templateCache[templateName]
	if !inMap {
		// need to create the template
		log.Println("creating template " + templateName + " and adding to cache")
		err = createTemplateCache(templateName)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have template in the cache
		log.Println("using cached template")
	}

	tmpl = templateCache[templateName]

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing template:", err)
	}
}

func createTemplateCache(templateName string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", templateName),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	parsedTemplate, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache
	templateCache[templateName] = parsedTemplate

	return nil
}
