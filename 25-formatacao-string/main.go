package main

import (
	"fmt"
	"strings"
)

func main() {
	name := " Paulo Henrique "
	fmt.Println("Original:", name)
	fmt.Println("Sem espaços:", strings.TrimSpace(name))
	fmt.Println("Em maiusculas:", strings.ToUpper(name))
	fmt.Println("Em minusculas:", strings.ToLower(name))
	fmt.Println("Substituição:", strings.ReplaceAll(name, "Henrique", "Reis"))

	name2 := "Maria"
	age := 28
	balance := 12998.56

	msg := fmt.Sprintf("Nome: %s\nIdade: %d\nSaldo: %.2f", name2, age, balance)
	fmt.Println(msg)

}
