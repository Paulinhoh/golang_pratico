package main

import "fmt"

type Mamiferos interface {
	saudacao()
	hobbie()
}

type Humano struct {
	Nome      string
	Profissao string
}

func (h Humano) saudacao() {
	fmt.Println("Oi meu nome é " + h.Nome + " e eu sou " + h.Profissao)
}

func (h Humano) hobbie() {
	fmt.Println("Video game")
}

type Cachorro struct {
	Nome string
}

func (c Cachorro) saudacao() {
	fmt.Println(c.Nome + " diz: au au")
}

func (c Cachorro) hobbie() {
	fmt.Println("brincar")
}

type Gato struct {
	Nome        string
	TempoDeSono int
}

func (g Gato) saudacao() {
	fmt.Println(g.Nome + " diz: miau")
}

func (g Gato) hobbie() {
	fmt.Printf("durmo uma %d horas\n", g.TempoDeSono)
}

func main() {
	paulinho := Humano{
		Nome:      "paulo h.",
		Profissao: "engenheiro de software",
	}

	luffy := Cachorro{
		Nome: "luffy",
	}

	spike := Gato{
		Nome:        "spike",
		TempoDeSono: 11,
	}

	apresentacao(paulinho)
	apresentacao(luffy)
	apresentacao(spike)
}

func apresentacao(m Mamiferos) {
	m.saudacao()
	m.hobbie()
}
