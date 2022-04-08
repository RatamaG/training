package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ingresa un numero ")

		return
	}
	name := os.Args[1]
	dim, err := strconv.Atoi(name)
	if err != nil {
		fmt.Println("ingresa un numero, letras no son compatibles")
	}
	for i := 0; i < dim; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	for i := 0; i < dim-1; i++ {
		fmt.Print("*")
		for i := 0; i < dim-2; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("*" + "\n")
	}
		


}
