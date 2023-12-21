package main

import "log"

func main() {
	myVar := "dog"

	switch myVar {
	case "cat":
		log.Println("Cat is set to cat")

	case "dog":
		log.Println("Cat is set to dog")

	case "fish":
		log.Println("Cat is set to fish")

	default:
		log.Println("None of them!")
	}
}
