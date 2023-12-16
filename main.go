package main

import "fmt"

// Função que adiciona um valor a uma variável usando um ponteiro
func addValue(num *int, value int) {
	*num += value
}

func main() {
	number := 10
	fmt.Println("Número antes da função:", number)

	// Passando o endereço da variável para a função usando &
	addValue(&number, 5)
	fmt.Println("Número após a função:", number)
}
