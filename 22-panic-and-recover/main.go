package main

import "fmt"

func safeOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperando do panic:", r)
		}
	}()
	panic("algo muito ruim aconteceu!!!")
}

func main() {
	fmt.Println("programa foi iniciado...")
	safeOperation()
	fmt.Println("programa finalizado!")
}
