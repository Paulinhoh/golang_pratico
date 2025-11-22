package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// json com structs forma mais correta de se utilizar json
type Employee struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Position string   `json:"position"`
	Skills   []string `json:"skills"`
	Address  Address  `json:"address"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func validateEmployee(e Employee) error {
	if e.Name == "" {
		return errors.New("o nome do funcionario é obrigatório")
	}
	if e.Position == "" {
		return errors.New("o cargo do funcionario é obrigatório")
	}
	return nil
}

func main() {
	jsonData := `{
		"id": 1,
		"name": "paulinho",
		"position": "software engineer",
		"skills": ["go", "python"],
		"address": {
			"city": "aracaju",
			"country": "brasil"
		}
	}`

	var employee Employee
	err := json.Unmarshal([]byte(jsonData), &employee)
	if err != nil {
		fmt.Println("falha ao decodificar json:", err)
		return
	}

	err = validateEmployee(employee)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("funcionario:", employee)
}
