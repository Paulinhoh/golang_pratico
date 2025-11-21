package main

import "fmt"

func main() {
	// na maioria dos programas em go o slice é muito mais usado que arrays por conta da sua flexibilidade

	var numbers []int // slice
	fmt.Println(numbers)

	days := []string{"segunda", "terça"}
	fmt.Println(days)
	fmt.Println()

	// append()
	numbers2 := []int{78, 90, 23}
	fmt.Println(numbers2)
	numbers2 = append(numbers2, 60, 90, 34, 21)
	fmt.Println(numbers2)
	fmt.Println()

	// subslice
	numbers3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers3Subslice := numbers3[3:6]
	fmt.Println(numbers3Subslice)

	numbers3Sub1 := numbers3[:2]
	fmt.Println(numbers3Sub1)
	numbers3Sub2 := numbers3[3:]
	fmt.Println(numbers3Sub2)
	fmt.Println()

	// Deletando uma posição atraves de subslices (deletando o valor 3)
	numbers3Delete3Index := append(numbers3[:2], numbers3[3:]...)
	fmt.Println(numbers3Delete3Index)
	fmt.Println()

	// matriz
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matriz)
	fmt.Println()

	// make (no make se pode definir capacidade e tamanho)
	// tamanho é o tamnho predefinido nesse caso 5
	// capacidade é até quando esse slice pode crescer, quando omitido o valor padrão é o mesmo do tamanho
	numbers4 := make([]int, 5)
	fmt.Println(numbers4)
	numbers4[0] = 9
	fmt.Println(numbers4)

	fmt.Println(len(numbers4))
	fmt.Println(cap(numbers4))
	fmt.Println()

	numbers5 := make([]int, 5, 10)
	fmt.Println(numbers5)

	for i := 0; i < 6; i++ {
		numbers5 = append(numbers5, i)
	}
	fmt.Println(numbers5)
	fmt.Println(len(numbers5))
	fmt.Println(cap(numbers5))
	// quando a capacidade é ultrapassada o go ele recalcula e dobra a capacidade do slice

}
