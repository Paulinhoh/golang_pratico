package main

import "fmt"

func main() {
	// operadores aritmeticos
	sub := 80 - 6
	sum := 40 + 7
	div := 46 / 7    //divisão inteira
	div2 := 46.0 / 7 //divisão (pelo menos um dos valores deve estar em float)
	mul := 56 * 23
	rest := 11 % 3

	fmt.Println(40 + 7)

	fmt.Println(sub)
	fmt.Println(sum)
	fmt.Println(div)
	fmt.Println(div2)
	fmt.Println(mul)
	fmt.Println(rest)

	// operadores relacionais
	fmt.Println(34 > 4)
	fmt.Println(34 < 4)
	fmt.Println(34 >= 4)
	fmt.Println(34 <= 4)
	fmt.Println(34 == 4)
	fmt.Println(34 != 4)

	// operadores de atribuição
	x := 10
	x += 5 // * / - + 5
	x++    // incrementa +1
	fmt.Println(x)

	// operadores lógicos &&(e) ||(ou) ~(não)
	userPass := "1234"
	isAdmin := true
	if userPass == "12345" || isAdmin {
		fmt.Println("usuario logado")
	} else {
		fmt.Println("não permitido")
	}

	// operadores bit a bit
	/*
		& (and)
		| (or)
		^ (xor)
		<< deslocamento para esquerda
		>> deslocamento para a direita
	*/
	a := 10 // 1010
	b := 3  // 0011
	result := a & b
	fmt.Println(result)
}
