package main

import "fmt"

// <-chan (somente leitura)
func readMessage(ch <-chan string) {
	msg1 := <-ch
	msg2 := <-ch
	fmt.Println("Mensagens recebidas:", msg1, "+", msg2)
}

func main() {
	ch := make(chan string, 2)

	ch <- "Message 1"
	ch <- "Message 2"

	readMessage(ch)
}
