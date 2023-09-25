package main

import "fmt"

func main() {
	fmt.Println("Hey !")

	var whatToSay string
	var i int

	whatToSay = "Goodbay"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid := saySomething()
	fmt.Println("The function returns", whatWasSaid)

	oneStringSaid, secondStringSaid := saySomethingMore()
	fmt.Println("Other function returns", oneStringSaid, secondStringSaid)
}

func saySomething() string {
	return "something"
}

func saySomethingMore() (string, string) {
	return "something", "ellse"
}
