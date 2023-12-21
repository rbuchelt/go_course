package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	var users []User
	users = append(users, User{"Renan", "Buchelt", "renan.oliveira@ituran.com.br", 36})
	users = append(users, User{"Priscilla", "Mendonça", "pri@ituran.com.br", 34})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
