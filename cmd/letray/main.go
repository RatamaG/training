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

	x := dim / 2

	for i := 0; i <= x-1; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}

		fmt.Print("*")

		for j := 1; j <= x-i-1; j++ {
			fmt.Print("  ")
		}

		fmt.Print("*")

		for j := dim - i; j < x; j++ {
			fmt.Print(" ")
		}

		fmt.Println()
	}
	if dim%2 == 0 {
		
		for i := 0; i < x; i++ {

			for i := 0; i < x-1; i++ {

				fmt.Print(" ")

			}

			fmt.Println("**")
		}
	} else {
		for i := 0; i < x+1; i++ {

			for i := 0; i < x-1; i++ {

				fmt.Print(" ")

			}

			fmt.Println("**")
		}
	}
}
