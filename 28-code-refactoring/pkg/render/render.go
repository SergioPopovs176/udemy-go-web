package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	fmt.Println("Try render template", tmpl)

	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	// fmt.Println(parsedTemplate)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing template:", err)
	}
}
