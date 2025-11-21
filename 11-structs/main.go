package main

import "fmt"

// struct
type Person struct {
	Name   string
	Age    int
	Email  string
	Adress Adress
	Employment
}

type Adress struct {
	Street string
	City   string
	Zip    string
}

type Employment struct {
	Position string
}

// método de Person
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func NewPerson(name, email, street, city, zip, position string, age int) Person {
	person := Person{
		Name:  name,
		Age:   age,
		Email: email,
		Adress: Adress{
			Street: street,
			City:   city,
			Zip:    zip,
		},
		Employment: Employment{
			Position: position,
		},
	}

	return person
}

func main() {
	pessoa := Person{
		Name:  "Paulo H",
		Age:   24,
		Email: "email@gmail.com",
		Adress: Adress{
			Street: "panificador silva",
			City:   "Aracaju",
			Zip:    "49107-240",
		},
	}
	pessoa.Position = "software engineer"

	fmt.Println(pessoa)
	fmt.Println("Name:", pessoa.Name)
	fmt.Println("Age:", pessoa.Age)
	fmt.Println("Email:", pessoa.Email)
	fmt.Println("Adress:", pessoa.Adress)
	fmt.Println("Adress street:", pessoa.Adress.Street)
	fmt.Println("Adress city:", pessoa.Adress.City)
	fmt.Println("Adress zip:", pessoa.Adress.Zip)
	fmt.Println("Position:", pessoa.Position)

	greet := pessoa.Greet()
	fmt.Println(greet)
	fmt.Println()

	pessoa1 := Person{
		Name:  "Maria",
		Age:   27,
		Email: "email@gmail.com",
		Adress: Adress{
			Street: "panificador ssouza",
			City:   "Baiha",
			Zip:    "48415-030",
		},
		Employment: Employment{
			Position: "data engineer",
		},
	}
	fmt.Println(pessoa1)
	fmt.Println()

	persons := []Person{}
	persons = append(persons, pessoa)
	persons = append(persons, pessoa1)
	fmt.Println(persons)

	for _, p := range persons {
		println(p.Greet())
	}

	pessoa2 := NewPerson("gabriel", "email@gamil.com", "fazenda laje da boa vista", "fatima", "48415-000", "Estudante", 13)
	fmt.Println(pessoa2)

}
