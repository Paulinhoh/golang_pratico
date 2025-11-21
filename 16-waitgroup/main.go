package main

import (
	"fmt"
	"sync"
	"time"
)

func makeCoffe(wg *sync.WaitGroup) {
	// defer -> garante que essa linha sera a ultima a ser rodada dentro da função
	defer wg.Done()

	fmt.Println("Fazendo o café...")
	time.Sleep(2 * time.Second)
	fmt.Println("Café pronto!")
}

func toastBread(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Fazendo a torrada...")
	time.Sleep(3 * time.Second)
	fmt.Println("Torrada pronta!")
}

func fryEggs(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Fritando ovos...")
	time.Sleep(1 * time.Second)
	fmt.Println("Ovos prontos!")
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(3)

	go makeCoffe(&wg)
	go toastBread(&wg)
	go fryEggs(&wg)

	// waitGroup esperara terminar os processos das funções na qual ele foi passado (nesse caso 3)
	wg.Wait()

	fmt.Println("Café da manha esta pronto!")
	fmt.Printf("Tempo total: %v\n", time.Since(start))
}
