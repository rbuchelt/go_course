package main

import "log"

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	me := User{
		FirstName: "Renan",
		LastName:  "Buchelt de Oliveira",
		Age:       36,
	}

	myMap := make(map[string]User)

	myMap["me"] = me

	log.Println(myMap["me"].Age)
}
