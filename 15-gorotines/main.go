package main

import (
	"fmt"
	"time"
)

func makeCoffe() {
	fmt.Println("Fazendo o café...")
	time.Sleep(2 * time.Second)
	fmt.Println("Café pronto!")
}

func toastBread() {
	fmt.Println("Fazendo a torrada...")
	time.Sleep(3 * time.Second)
	fmt.Println("Torrada pronta!")
}

func fryEggs() {
	fmt.Println("Fritando ovos...")
	time.Sleep(1 * time.Second)
	fmt.Println("Ovos prontos!")
}

// a função main ja é uma go routine
func main() {
	start := time.Now()
	// rodando de maneira concorrente
	go makeCoffe()
	go toastBread()
	go fryEggs()

	time.Sleep(4 * time.Second)

	fmt.Println("Café da manha esta pronto!")
	fmt.Printf("Tempo total: %v\n", time.Since(start))
}
