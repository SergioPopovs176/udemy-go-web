package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString)

	log.Println("After myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)

	newValue := "Red"

	*s = newValue
}
