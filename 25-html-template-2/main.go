package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// ctrl+C to stop programm

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/links", Links)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	// n, err := fmt.Fprintf(w, "This is the HOME page")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(fmt.Sprintf("Bytes written: %d", n))

	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Header.Values("User-Agent"))

	// n, err := fmt.Fprintf(w, "This is the ABOUT page")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(fmt.Sprintf("Bytes written: %d", n))

	renderTemplate(w, "about.page.tmpl")
}

func Links(w http.ResponseWriter, r *http.Request) {
	// n, err := fmt.Fprintf(w, "This is the LINKS page")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	renderTemplate(w, "links.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
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
