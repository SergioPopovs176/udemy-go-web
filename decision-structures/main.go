package main

import "log"

func main() {
	// var isTrue bool
	// isTrue = true
	isTrue := true
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 100
	isTrue2 := false
	if myNum > 99 && !isTrue2 {
		log.Println("myNum is greate 99 and not true")
	} else if myNum < 100 && isTrue2 {
		log.Println("myNum is less 100 and true")
	} else if myNum == 101 || isTrue2 {
		log.Println("myNum equel 101 or true")
	} else if myNum > 1000 && isTrue == false {
		log.Println("myNum is greate 1000 and not true")
	}

	myVar := "dogs"
	// Если сработает какое-то условие то switch заканчивается и остальные условия проверяться не будут !
	switch myVar {
	case "cat":
		log.Println("Var is set to cat")
	case "dog":
		log.Println("Var is set to dog")
	case "fish":
		log.Println("Var is set to fish")
	default:
		log.Println("Var is set to", myVar)
	}
}
