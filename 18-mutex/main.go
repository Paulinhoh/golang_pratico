package main

import (
	"fmt"
	"sync"
)

var (
	counter int        // recurso compartilhado
	mutex   sync.Mutex // tratar condição de corrida (race condicional)
)

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock() // travando para garantir que só uma go routine acesse a variavel couter por vez
	counter++
	mutex.Unlock() // destravando o acesso
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10_000)
	for range 10_000 {
		go incrementCounter(&wg)
	}

	wg.Wait()
	fmt.Println(counter)
}
