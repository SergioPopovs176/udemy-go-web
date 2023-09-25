package main

// ctrl+C остановить выполнение программы
// сделать программу модульной
// go mod init mywebapp
import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// Home is the about page handler
func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is home page")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Number of bytes written : %d", n))
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	n, err := fmt.Fprintf(w, fmt.Sprintf("This is about page and 2+2 is %d", sum))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Number of bytes written : %d", n))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x float32
	var y float32

	x = 23.56
	y = 2.34

	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

// addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	return x / y, nil
}
