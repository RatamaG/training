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

	for i := 0; i < dim/2-2; i++ {

		fmt.Print("*")

		for i := 0; i < dim-2; i++ {

			fmt.Print(" ")
		}

		fmt.Print("*" + "\n")
	}

	for i := 0; i < dim; i++ {

		fmt.Print("*")

	}

	fmt.Println()

	for i := 0; i < dim/2-2; i++ {

		for i := 0; i < dim-1; i++ {

			fmt.Print(" ")

		}

		fmt.Print("*")
		fmt.Println()

	}

	for i := 0; i < dim/2; i++ {
		fmt.Print(" ")
	}
	for i := 0; i < dim; i++ {

		fmt.Print("*")

	}

	fmt.Println()

	for i := 0; i < dim/2-2; i++ {

		for i := 0; i < dim-1; i++ {

			fmt.Print(" ")

		}

		fmt.Print("*")
		fmt.Println()

	}
}
