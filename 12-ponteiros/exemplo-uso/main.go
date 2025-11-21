package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) showName() {
	fmt.Println(&p.name)
}

func main() {
	age := 25
	fmt.Println(&age) // endereço age da main
	fmt.Println()

	showAgeExemple(age)

	showAge(&age)
	fmt.Println(age)
	fmt.Println()

	// -=-=-=-=-===-=-=-=-=-=-=-=-=-=-=-=
	// ponteiros com metodos
	p := Person{
		name: "paulinho",
	}

	fmt.Println(&p.name)
	p.showName()
}

func showAgeExemple(age int) {
	// quando a variavel age foi chamada na função o go faz uma cópia do valor
	fmt.Println(&age) // endereço age de showAge
	fmt.Println(age)
	fmt.Println()
}

// passagem por valor = age
// passagem por referência = age *int
func showAge(age *int) {
	fmt.Println(age)
	fmt.Println(*age)
	*age = 30
	fmt.Println()
}
