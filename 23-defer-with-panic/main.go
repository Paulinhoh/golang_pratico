package main

import (
	"fmt"
	"os"
)

func writeFile() {
	file, err := os.Create("23-defer-with-panic/example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintln(file, "hello, world!")
	fmt.Println("Arquivo escrito corretamente.")
}

func main() {
	writeFile()
}
