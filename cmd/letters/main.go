package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Only two arguments are allowed, a number in the first position and a letter in the second.")

		return
	}
	if len(os.Args) > 3 {
		fmt.Println("Error: Only two arguments are allowed, a number in the first position and a letter in the second.")

		return
	}

	name := os.Args[1]
	dim, err := strconv.Atoi(name)

	if err != nil {
		fmt.Print(" Error : Enter a number in the first position and a letter in the second ")
	}

	letter := os.Args[2]
	
	let := strings.ToLower(letter)

	switch let {

	case "a":

		//lettera
		a := "*"

		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}

		fmt.Println()

		for i := 0; i < dim/2-2; i++ {
			fmt.Print(a)

			for i := 0; i < dim-2; i++ {
				fmt.Print(" ")
			}

			fmt.Print(a + "\n")
		}

		for i := 0; i < dim; i++ {
			fmt.Print(a)
		}

		fmt.Println()

		for i := 0; i < dim/2; i++ {
			fmt.Print(a)
			for i := 0; i < dim-2; i++ {
				fmt.Print(" ")
			}

			fmt.Print(a + "\n")
		}

		if dim%2 == 0 {
		} else {
			for i := 0; i < dim-dim+1; i++ {
				fmt.Print(a)
				for i := 0; i < dim-2; i++ {
					fmt.Print(" ")
				}

				fmt.Print(a + "\n")
			}
		}
		//lettera

	case "b":

		//letterb
		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}

		fmt.Println()

		for i := 0; i < dim/2-1; i++ {
			fmt.Print("*")
			for i := 0; i < dim-2; i++ {
				fmt.Print(" ")
			}

			fmt.Print("*" + "\n")
		}

		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}

		fmt.Println("")

		if dim%2 == 0 {

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

			fmt.Println("")

		} else {

			for i := 0; i < dim/2-1; i++ {
				fmt.Print("*")
				for i := 0; i < dim-2; i++ {
					fmt.Print(" ")
				}

				fmt.Print("*" + "\n")
			}

			for i := 0; i < dim; i++ {
				fmt.Print("*")
			}

			fmt.Println("")
		}
		//letterb

	case "c":

		//letterc
		for i := 0; i < dim-1; i++ {
			fmt.Print("*")
		}

		for i := 0; i < dim-1; i++ {
			fmt.Println("*")
		}

		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}

		fmt.Println("")
		//letterc

	case "d":

		//letterd
		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}

		fmt.Println()

		for i := 0; i < dim-2; i++ {

			fmt.Print("*")

			for i := 0; i < dim-2; i++ {

				fmt.Print(" ")
			}

			fmt.Print("*" + "\n")
		}

		for i := 0; i < dim; i++ {

			fmt.Print("*")

		}

		fmt.Println("")
		//letterd

	case "e":

		//lettere
		for i := 0; i < dim-1; i++ {
			fmt.Print("*")
		}

		for i := 0; i < dim/2; i++ {
			fmt.Println("*")
		}

		for i := 0; i < dim-1; i++ {
			fmt.Print("*")
		}

		if dim%2 == 0 {
			for i := 0; i < dim/2-1; i++ {
				fmt.Println("*")
			}
		} else {

			for i := 0; i < dim/2; i++ {
				fmt.Println("*")

			}
		}

		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}

		fmt.Println("")
		//lettere

	case "f":

		// letterf
		for i := 0; i < dim-1; i++ {
			fmt.Print("*")
		}

		for i := 0; i < dim/2; i++ {
			fmt.Println("*")
		}

		for i := 0; i < dim-1; i++ {
			fmt.Print("*")
		}

		if dim%2 == 0 {
			for i := 0; i < dim/2; i++ {
				fmt.Println("*")
			}
		} else {

			for i := 0; i < dim/2+1; i++ {

				fmt.Println("*")

			}
		}
		// letterf

	case "g":

		// letterg
		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}

		fmt.Println()

		for i := 0; i < dim/2-1; i++ {

			fmt.Print("*")

			for i := 0; i < dim-2; i++ {
				fmt.Print(" ")
			}

			fmt.Print("*" + "\n")
		}

		for i := 0; i < dim; i++ {

			fmt.Print("*")
		}

		fmt.Println("")

		if dim%2 == 0 {

			for i := 0; i < dim/2-2; i++ {

				for i := 0; i < dim-1; i++ {

					fmt.Print(" ")

				}

				fmt.Print("*")
				fmt.Println()
			}
			for i := 0; i < dim; i++ {
				fmt.Print("*")
			}

		} else {
			for i := 0; i < dim/2-1; i++ {
				for i := 0; i < dim-1; i++ {

					fmt.Print(" ")
				}

				fmt.Print("*")
				fmt.Println()
			}

			for i := 0; i < dim; i++ {
				fmt.Print("*")
			}
		}
		// letterg

	case "h":

		// letterh
		for i := 0; i < dim/2; i++ {

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

		for i := 0; i < dim/2-1; i++ {
			fmt.Print("*")
			for i := 0; i < dim-2; i++ {
				fmt.Print(" ")
			}

			fmt.Print("*" + "\n")
		}

		if dim%2 == 0 {
		} else {
			for i := 0; i < dim-dim+1; i++ {
				fmt.Print("*")
				for i := 0; i < dim-2; i++ {
					fmt.Print(" ")
				}

				fmt.Print("*" + "\n")
			}
		}
		// letterh

	case "i":

		// letteri
		if err != nil {
			fmt.Println("enter a number, letters are not allowed")
		}

		for i := 0; i < dim; i++ {

			fmt.Println("*")

		}
		// letteri

	case "j":
		// letterj
		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}

		fmt.Println()

		for i := 0; i < dim/2-1; i++ {

			for i := 0; i < dim-1; i++ {

				fmt.Print(" ")
			}

			fmt.Print("*" + "\n")
		}
		//
		if dim%2 == 0 {

			for i := 0; i < dim/2-1; i++ {

				fmt.Print("*")

				for i := 0; i < dim-2; i++ {
					fmt.Print(" ")
				}

				fmt.Print("*" + "\n")
			}

			for i := 0; i < dim; i++ {

				fmt.Print("*")
			}
		} else {
			for i := 0; i < dim/2; i++ {

				fmt.Print("*")

				for i := 0; i < dim-2; i++ {
					fmt.Print(" ")
				}

				fmt.Print("*" + "\n")
			}

			for i := 0; i < dim; i++ {

				fmt.Print("*")
			}
		}
		// letterj

	case "k":

		// letterk
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
		// letterk

	case "l":

		// letterl
		for i := 0; i < dim-1; i++ {

			fmt.Println("*")

		}

		for i := 0; i < dim; i++ {

			fmt.Print("*")

		}
		// letterl

	case "m":

		// letterm
		for i := 0; i <= dim-1; i++ {

			fmt.Print("*")

			for j := 1; j <= i; j++ {
				fmt.Print(" ")
			}

			fmt.Print("*")

			for j := 1; j <= dim-i-1; j++ {
				fmt.Print("  ")
			}

			fmt.Print("*")

			for j := dim - i; j < dim; j++ {
				fmt.Print(" ")
			}

			fmt.Print("*")
			fmt.Println()
		}
		// letterm

	case "n":

		// lettern
		for i := 0; i < dim; i++ {

			fmt.Print("*")

			for j := 1; j <= i; j++ {
				fmt.Print(" ")
			}

			fmt.Print("*")

			for j := 1; j <= dim-i; j++ {
				fmt.Print(" ")
			}

			fmt.Print("*")
			fmt.Println()
		}
		// lettern

	case "o":

		// lettero
		for i := 0; i < dim; i++ {

			fmt.Print("*")

		}

		fmt.Println()

		for i := 0; i < dim-2; i++ {

			fmt.Print("*")

			for i := 0; i < dim-2; i++ {
				fmt.Print(" ")
			}

			fmt.Print("*" + "\n")
		}

		for i := 0; i < dim; i++ {

			fmt.Print("*")
		}

		fmt.Println("")
		// lettero

	case "p":

		// letterp
		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}

		fmt.Println()

		for i := 0; i < dim/2-1; i++ {

			fmt.Print("*")

			for i := 0; i < dim-2; i++ {

				fmt.Print(" ")
			}

			fmt.Print("*" + "\n")
		}

		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}
		if dim%2 == 0 {

			fmt.Println()

			for i := 0; i < dim/2-1; i++ {

				fmt.Println("*")

			}

			fmt.Println("")
		} else {

			fmt.Println()

			for i := 0; i < dim/2; i++ {

				fmt.Println("*")

			}
		}
		// letterp

	case "q":

		// letterq
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
		// letterq

	case "r":

		// letterr
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
		if dim%2 == 0 {

			fmt.Println()

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
		} else {

			fmt.Println()

			for i := 0; i < dim/2+1; i++ {

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
		}
		// letterr

	case "s":

		// letters

		for i := 0; i < dim-1; i++ {
			fmt.Print("*")
		}

		for i := 0; i < dim/2; i++ {
			fmt.Println("*")
		}

		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}

		fmt.Println()

		for i := 0; i < dim/2-1; i++ {

			for i := 0; i < dim-1; i++ {
				fmt.Print(" ")
			}
			fmt.Print("*")
			fmt.Println()
		}

		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}
		// letters

	case "t":

		// lettert
		for i := 0; i < dim; i++ {

			fmt.Print("*")

		}
		fmt.Println()

		for i := 0; i < dim-1; i++ {

			for i := 0; i < dim/2; i++ {

				fmt.Print(" ")

			}

			fmt.Print("*")
			fmt.Println()
		}
		// lettert

	case "u":

		// letteru
		for i := 0; i < dim-1; i++ {

			fmt.Print("*")

			for i := 0; i < dim-2; i++ {
				fmt.Print(" ")
			}

			fmt.Print("*" + "\n")
		}

		for i := 0; i < dim; i++ {

			fmt.Print("*")
		}

		fmt.Println("")
		// letteru

	case "v":

		//letteru
		for i := 0; i <= dim-1; i++ {

			for j := 1; j <= i; j++ {

				fmt.Print(" ")

			}

			fmt.Print("*")

			for j := 1; j <= dim-i-1; j++ {

				fmt.Print("  ")

			}

			fmt.Print("*")
			fmt.Println()
		}
		//letteru

	case "w":

		//letterw
		for i := 0; i <= dim-1; i++ {

			fmt.Print("*")

			for j := 1; j <= dim-i; j++ {

				fmt.Print(" ")

			}

			fmt.Print("*")

			for j := 1; j <= i; j++ {

				fmt.Print("  ")

			}

			fmt.Print("*")

			for j := 0; j < dim-i; j++ {

				fmt.Print(" ")

			}

			fmt.Print("*")
			fmt.Println()
		}
		//letterw

	case "x":

		//letterx
		for i := dim / 2; i <= dim-1; i++ {
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
		//Second part
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
		//letterx

	case "y":

		//lettery
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
		//lettery

	case "z":

		// letterz
		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}

		fmt.Println()

		for i := 0; i < dim-2; i++ {

			for j := 1; j <= dim-i-2; j++ {
				fmt.Print(" ")
			}

			fmt.Print("*")
			fmt.Println()
		}

		for i := 0; i < dim; i++ {
			fmt.Print("*")
		}
		// letterz
	}
}
