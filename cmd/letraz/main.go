package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("enter a number")

		return
	}
	
	name := os.Args[1]
	dim, err := strconv.Atoi(name)
	if err != nil {
		fmt.Println("enter a number, letters are not allowed")
	}

	for i := 0; i < dim; i++ {
		fmt.Print("*")
	}

	fmt.Println()
	
	for i := 0; i < dim-2; i++ {

		for j :=1; j <= dim-i-2; j++ {
			fmt.Print(" ")
		}

		fmt.Print("*")
		fmt.Println()
	}

	for i := 0; i < dim; i++ {
		fmt.Print("*")
	}

}
