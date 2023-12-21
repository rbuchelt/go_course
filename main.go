package main

import "fmt"

func main() {
	// Comunicação Síncrona
	canalSincrono := make(chan string)

	// Goroutine enviando dados para o canal
	go func() {
		canalSincrono <- "Mensagem Síncrona"
	}()

	// Recebendo dados (bloqueia até receber)
	mensagem := <-canalSincrono
	fmt.Println(mensagem) // Saída: Mensagem Síncrona

	// Comunicação Assíncrona
	canalAssincrono := make(chan string)

	// Goroutine enviando dados para o canal
	go func() {
		canalAssincrono <- "Mensagem Assíncrona"
	}()

	// Recebendo dados (não bloqueia)
	select {
	case mensagem := <-canalAssincrono:
		fmt.Println(mensagem) // Saída: Mensagem Assíncrona
	default:
		fmt.Println("Nenhuma mensagem disponível")
	}
}
