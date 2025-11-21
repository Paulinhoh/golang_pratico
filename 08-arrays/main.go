package main

import "fmt"

func main() {
	// forma basica
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 3
	numbers[2] = 67
	fmt.Println(numbers)
	
	// iniciando ja com valores
	days := [3]string{"segunda", "terça", "quarta"}
	fmt.Println(days)

	// Arrays não se pode aumentar a capacidade
	// iniciando sem dar o valor da tamanho, mas ele pega o valor do tamanho pela quantidade de elementos quando iniciado, no exemplo abaixo [3]
	notas := [...]int{55, 23, 67}
	fmt.Println(notas)
}
