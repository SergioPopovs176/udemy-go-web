package main

import (
	"log"
	"time"
)

var s = "seven"

var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time

// Вместо много переменных поместим всё в структуру

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	var s2 = "six"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")

	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawer",
		PhoneNumber: "555-678-000",
		Age:         34,
	}

	log.Println(user.FirstName, user.LastName, "Phone:", user.PhoneNumber, "Age:", user.Age, "Birthday:", user.BirthDate)
}

func saySomething(s string) (string, string) {
	log.Println("s from the saySomething func is", s)

	return s, "World"
}
