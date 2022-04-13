package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("enter a number ")

		return
	}

	name := os.Args[1]
	dim, err := strconv.Atoi(name)
	if err != nil {
		fmt.Println("enter a number, letters are not allowed")
	}

	a := "*"
	//First module
	for i := 0; i < dim; i++ {
		fmt.Print("*")
	}

	fmt.Println()

	//Second module
	for i := 0; i < dim/2-2; i++ {
		fmt.Print(a)
		//dim
		for i := 0; i < dim-2; i++ {
			fmt.Print(" ")
		}

		fmt.Printf(a + "\n")
	}

	for i := 0; i < dim; i++ {
		fmt.Print(a)
	}

	//Third module
	fmt.Println()
	for i := 0; i < dim/2; i++ {
		fmt.Print(a)
		for i := 0; i < dim-2; i++ {
			fmt.Print(" ")
		}

		fmt.Printf(a + "\n")
	}

	if dim%2 == 0 {
	} else {
		for i := 0; i < dim-dim+1; i++ {
			fmt.Print(a)
			for i := 0; i < dim-2; i++ {
				fmt.Print(" ")
			}
			
			fmt.Printf(a + "\n")
		}
	}
}
