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
	sb := 0
	for i := 0; i < dim-1; i++ {
		for j := 0; j < sb; j++ {
			fmt.Print(" ")
		}

		fmt.Print("*")
		sm := dim - (i+1)*2
		if sm < 0 {
			sm = sm * -1
		}
		for j := 0; j < sm; j++ {
			fmt.Print(" ")
		}
		fmt.Print("*", "\n")
		if i >= ((dim / 2) - 1) {
			sb--
		} else {
			sb++
		}
	}
}