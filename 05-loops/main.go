package main

import "fmt"

func main() {
	// loop for - basico
	for i := 0; i < 5; i++ {
		fmt.Println("Valor de i: ", i)
	}
	fmt.Println()

	// for como while
	i := 0
	for i < 5 {
		fmt.Println("Valor de i: ", i)
		i++
	}
	fmt.Println()

	// loop infinito
	i = 0
	for {
		fmt.Println("loop infinito")
		if i > 3 {
			break
		}
		i++
	}
	fmt.Println()

	// for com range
	numbers := []int{10, 20, 30} // slice
	for index, value := range numbers {
		fmt.Printf("indice: %d, valor: %d\n", index, value)
	}
	fmt.Println()

	// caso voce por exemplo não quisesse pegar o index basta omiti-lo
	for _, value := range numbers {
		fmt.Println("Valor: ", value)
	}
	fmt.Println()

	capitais := map[string]string{ // map
		"Brasil": "Brasilia",
		"França": "Paris",
		"Japão":  "Tokyo",
	}
	for key, value := range capitais {
		fmt.Printf("pais: %s, capital: %s\n", key, value)
	}
}
