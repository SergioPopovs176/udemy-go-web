package main

import (
	"net/http"
)

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
