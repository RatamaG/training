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

	for i := 0; i <= dim/2-1; i++ {

		fmt.Print("*")

		for j := 1; j <= dim/2-i; j++ {
			fmt.Print(" ")
		}

		fmt.Print("*")
		fmt.Println()
	}
	//
	for i := 0; i < dim/2; i++ {

		fmt.Print("*")

		for j := 1; j <= i; j++ {

			fmt.Print("  ")

		}

		fmt.Print("*")

		for j := 1; j <= dim-i; j++ {

			fmt.Print("  ")

		}

		fmt.Println()
	}
	if dim%2 == 0 {
	} else {

		for i := 0; i < dim-dim+1; i++ {

			fmt.Print("*")

			for j := 0; j < dim-i-1; j++ {

				fmt.Print(" ")

			}

			fmt.Print("*" + "\n")
		}
	}

}
