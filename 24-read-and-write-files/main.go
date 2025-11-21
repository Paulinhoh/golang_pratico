package main

import (
	"bufio"
	"fmt"
	"os"
)

func simpleReadFile() {
	data, err := os.ReadFile("24-read-and-write-files/nomes.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}
	fmt.Println("Conteudo do arquivo:")
	println(string(data))
}

func lineReadFile() {
	file, err := os.Open("24-read-and-write-files/nomes.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Linhas do meu arquivo:")

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro durante a leitura:", err)
	}

}

func writeAlreadyExisteFile() {
	file, err := os.OpenFile("24-read-and-write-files/times.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo para escrita:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\nBahia")
	if err != nil {
		fmt.Println("Falha ao adicionar string ao arquivo:", err)
		return
	}
	fmt.Println("Dados adicionados com sucesso!")
}

func writeFile() {
	data := "BMW \nAudi i8"
	err := os.WriteFile("24-read-and-write-files/carros.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
	fmt.Println("Dados escritos no arquivo")
}

func main() {
	fmt.Println("Leitura simples")
	simpleReadFile()
	fmt.Println()

	fmt.Println("Leitura linha a linha")
	lineReadFile()
	fmt.Println()

	fmt.Println("Escrevendo em arquivo")
	writeFile()
	fmt.Println()

	fmt.Println("Escrevendo no fim de um arquivo existente")
	writeAlreadyExisteFile()
}
