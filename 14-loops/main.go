package main

import "log"

func main() {
	log.Println("----- SIMPLE LOOP -----")
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	log.Println("----- SLICE RANGE LOOP -----")
	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for i, animal := range animals {
		log.Println(i, animal)
	}
	for _, animal := range animals {
		log.Println(animal)
	}

	log.Println("----- MAP RANGE LOOP -----")
	animals2 := make(map[string]string)
	animals2["dog"] = "Fido"
	animals2["cat"] = "Fluffy"
	for animalType, animal := range animals2 {
		log.Println(animalType, animal)
	}
	for _, animal := range animals2 {
		log.Println(animal)
	}

	log.Println("----- STRING RANGE LOOP -----")
	var firstLine = "Once upon a midnight dreary"
	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	log.Println("----- CUSTOM RANGE LOOP -----")
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 21})
	users = append(users, User{"Sally", "Brown", "sally@brown.com", 35})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 12})
	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
