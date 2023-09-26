package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// MAP
	var myString string
	var myInt int

	myString = "Hi"
	myInt = 11

	mySecondString := "another string"

	log.Println(myString, mySecondString, myInt)

	myMap := make(map[string]string)

	myMap["dog"] = "Samon"
	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]int)

	myMap2["First"] = 1
	myMap2["Second"] = 2

	log.Println(myMap2["First"], myMap2["Second"])

	myMap3 := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMap3["me"] = me

	log.Println(myMap3["me"])
	log.Println(myMap3["me"].FirstName)

	me.FirstName = "fff"
	log.Println(myMap3["me"].FirstName)

	// SLICE
	var mySlice []string
	log.Println(mySlice)

	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	log.Println(mySlice)

	var mySlice2 []int
	log.Println(mySlice)

	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 4)
	log.Println(mySlice2)

	sort.Ints(mySlice2)
	log.Println(mySlice2)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)
	log.Println(numbers[0:2])

	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
