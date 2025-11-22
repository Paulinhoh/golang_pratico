package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{"name": "paulinho", "attributes":{"height": 1.8, "skills": ["go", "python"]}}`

	// interface para ser um tipo coringa (string, int, struct)
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("falha ao decodificar e json:", err)
		return
	}

	fmt.Println("Nome:", data["name"])
	attributes := data["attributes"].(map[string]interface{})
	fmt.Println("Altura:", attributes["height"])
	fmt.Println("Habilidades:", attributes["skills"])
}
