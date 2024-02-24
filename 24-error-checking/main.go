package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// ctrl+C to stop programm

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

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
	// fmt.Println(r.Header.Values("User-Agent"))

	n, err := fmt.Fprintf(w, "This is the ABOUT page")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Bytes written: %d", n))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 5.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")

		return
	}

	n, err := fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 5.0, f))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Bytes written: %d", n))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")

		return 0, err
	}

	result := x / y

	return result, nil
}
