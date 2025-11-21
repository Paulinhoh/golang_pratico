package main

import "fmt"

func main() {
	// & -> endereço - obtem o esndereço de uma variavel
	// * -> acessar o valor armazenado no endeteço que o ponteiro aponta

	name := "paulinho"
	var pointer *string

	fmt.Println(name)
	fmt.Println(pointer)
	fmt.Println()

	// endereços de memorias das variaveis
	fmt.Println(&name)
	fmt.Println(&pointer)
	fmt.Println()

	// poiter referenciando name
	pointer = &name
	fmt.Println(&pointer) // endereço da memoria de pointer
	fmt.Println(pointer)  // pointer apontando para a variavel name
	fmt.Println(*pointer) // pegando o valor da variavel que esta sendo apontanda (name)
	fmt.Println()

	// alteranção do valor de name pelo ponteiro
	*pointer = "paulo henrique"
	fmt.Println(*pointer)
	fmt.Println(name)
	fmt.Println()

	// alteração do valor de name pela propria variavel name
	*pointer = "paulo henrique"
	fmt.Println(*pointer)
	fmt.Println(name)
	fmt.Println()
}
