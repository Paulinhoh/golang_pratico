package main

import "fmt"

func main() {
	// if - else
	age := 17
	if age >= 18 && age < 80 {
		fmt.Println("maior de idade")
	} else if age >= 80 {
		fmt.Println("idoso")
	} else {
		fmt.Println("menor de idade")
	}

	// switch-case
	day := "sabadou"
	switch day {
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("dia util")
	case "sabado", "domingo":
		fmt.Println("final de semana")
	default:
		fmt.Println("dia desconhecido")
	}
}
