package main

import "fmt"

func main() {
	// função basica
	result := sum(3, 4)
	fmt.Println(result)

	// retornando multiplos valores
	_, remainder := divide(10, 3)
	// fmt.Println(quotient)
	fmt.Println(remainder)

	// função sem retorno
	hello("paulinho")
	fmt.Println()

	// =-=-=--==-=-=-=-=-=-=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-
	// funções anonimas
	anom := func(a, b int) int {
		return a + b
	}
	fmt.Println(anom(6, 8))

	// closures
	nextOdd := oddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	// funções variadicas
	fmt.Println(max(1, 8, 5, 90, 45, -122))
}

// funções
func sum(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

func hello(name string) {
	fmt.Println("hello,", name)
}

func oddGenerator() func() int {
	i := -1
	return func() int {
		i += 2
		return i
	}
}

func max(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	bigger := numbers[0]
	for _, num := range numbers {
		if num > bigger {
			bigger = num
		}
	}
	return bigger
}
