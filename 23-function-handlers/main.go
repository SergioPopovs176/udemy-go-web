package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// ctrl+C to stop programm

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is the HOME page")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Bytes written: %d", n))
}

func About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is the ABOUT page")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Bytes written: %d", n))
}
