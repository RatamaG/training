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
	//Primera parte
	for i := dim/2; i <= dim-1; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		for j := 1; j <= dim-i; j++ {
			fmt.Print("  ")
		}
		fmt.Print("*")
		fmt.Println()
	}
	//segnda parte
	for i := 0; i <= dim/2-1; i++ {
		for j := 1; j <= dim-i; j++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		for j := 1; j <= i; j++ {
			fmt.Print("  ")
		}
		fmt.Print("*")
		fmt.Println()
	}
}
