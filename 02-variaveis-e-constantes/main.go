package main

import "fmt"

// constantes
const Name4 = "constante: paulinho"

func main() {
	// Usando o var
	var name string

	// outra maneira já atribuindo o valor
	var name2 string = "Paulo H"

	// inferencia de tipo
	name3 := "Henrique"

	// printando os resultados
	name = "Paulinho"
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(Name4)

	// tipos de dados
	var age int = 28
	var pi float32 = 3.14  // 6-7 digitos de precisão
	var pi2 float64 = 3.14 //15-16 digitos de precisão
	var isEnabled bool = true

	/*
		- int8: representa numeros inteiros de 8 bits
			- intervalo: -128 a 127
		- int16:  representa numeros inteiros de 16 bits
			- intervalo: -32,768 a 32,767
		- int32: representa numeros inteiros de 32 bits
			- intervalo: -2,147,483,648 a 2,147,483,647
		- int64: representa numeros inteiros de 64 bits
			- intervalo: -9,223,372,036,854,775,808 a 9,223,372,036,854,775,807
		- int: representa numeros inteiros cujo tamanho da arquitetura (32bits ou 64bits)
			- intervalo: depende da arquitetura (em sistemas de 32 bits sera igual ao int32, se for 64 bits é igual a int64)
	*/

	/*
		inteiros sem sinal:
		- uint8: representa numeros inteiros sem sinal de 8 bits
			- intervalo: 0 a 255
		- uint16: representa numeros inteiros sem sinal de 16 bits
			- intervalo: 0 a 65,535
		- uint32: representa numeros inteiros sem sinal de 32 bits
			- intervalo: 0 a 4,294,967,295
		- uint64: representa numeros inteiros sem sinal de 64 bits
			- intervalo: 0 a 18,446,744,073,709,551,615
		- uint: representa numeros inteiros sem sinal cujo tamanho da arquitetura (32bits ou 64bits)
			- intervalo: depende da arquitetura (em sistemas de 32 bits, é igual a uint32, se for 64 bits, é igual a uint64)
	*/

	/*
		Outros tipos:
		- complex64: representa números complexos com partes reais e imaginarias do tipo float32
		- complex128: representa números complexos com partes reais e imaginarias do tipo float64
	*/

	fmt.Println(age)
	fmt.Println(pi)
	fmt.Println(pi2)
	fmt.Println(isEnabled)

	// multiplas variaveis
	var a, b, c = "a", 50, true
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

/*
	bloco de comentarios
*/
