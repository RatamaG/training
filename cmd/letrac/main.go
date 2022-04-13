package main

import (
	"fmt"
	"strconv"
	"os"
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
		//First modulo
		for i := 0; i < dim-1; i++ {
			fmt.Printf("*")
		}
		//Middle module
	
 		for i := 0; i < dim-1; i++ {
			fmt.Println("*")
		}
		//Last module
		
		for i := 0; i < dim; i++ {
			fmt.Printf("*")
		}
		
		fmt.Println("")
	
}