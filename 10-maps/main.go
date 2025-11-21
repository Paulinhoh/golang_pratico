package main

import (
	"fmt"
)

func main() {
	// modo de criação mais trabalhosa
	var m map[string]int     // criação do map
	m = make(map[string]int) // inicialização do map (todo map precisa ser inicializado)

	m["Paulinho"] = 24
	m["Ana Clara"] = 5
	fmt.Println(m)

	// recuperando um valor no map
	fmt.Println(m["Ana Clara"])
	fmt.Println(m["asdsfgrgsf"])

	age, exists := m["Paulinho"]
	if exists {
		fmt.Println("Paulinho existe no map")
		fmt.Println(age)
	}

	delete(m, "Paulinho")
	fmt.Println(m)
	fmt.Println()

	// forma mais comun de criar um map
	capitais := make(map[string]string) // criação e inicialização juntos
	capitais["são paulo"] = "são paulo"
	capitais["ceara"] = "fortaleza"
	capitais["pernambuco"] = "recife"
	fmt.Println(capitais)

	for estado, capital := range capitais {
		fmt.Println()
		fmt.Println(estado)
		fmt.Println(capital)
	}

	// declaração e inicialização com valores
	idade := map[string]int{"maria": 27}
	fmt.Println(idade)
	fmt.Println()

	// maps aninhados
	innerMaps := make(map[string]map[string]int)
	innerMaps["outer"] = make(map[string]int)
	innerMaps["outer"]["inner"] = 38
	fmt.Println(innerMaps)

	// tamanho
	size := make(map[string]string, 2)
	fmt.Println(size)
}
