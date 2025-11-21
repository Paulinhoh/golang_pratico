package main

import "fmt"

// chan<- (somente escrita)
func sendMessage(ch chan<- string) {
	ch <- "Message 01"
	ch <- "Message 02"
}
// somente chan (canal de leitura e escrita)

func main() {
	ch := make(chan string, 2)

	go sendMessage(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
