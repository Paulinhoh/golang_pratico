package main

import (
	"fmt"

	"github.com/Paulinhoh/golang_pratico/07-modulos/util"
	"github.com/Rhymond/go-money"
)

func main() {
	// chamando a soma no mesmo modulo
	fmt.Println(sum(3, 6))

	// chamar apartir de um modulo diferente
	fmt.Println(util.Sum(10, 10))
	fmt.Println()

	// Usando lib externa
	fmt.Println(6.50) // forma correta de apresentação seria R$6,50
	pound := money.NewFromFloat(6.50, money.BRL)
	fmt.Println(pound.Display())

}

// go run main.go (vai dar erro pois não execultou o arquivo math.go)
// go run . (execulta todos os arquivos do diretorio) (funciona)

// pkg.go.dev (site para procurar bibliotecas go)
// go get nome_biblioteca (comando para instalar a biblioteca)
