package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ingresa un nombre ")


		return
	}
	name := os.Args[1]
	dim, err := strconv.Atoi(name)
	if err != nil {
			fmt.Println("ingresa un numero, letras no son compatibles")
		} 
		//Primer modulo
		for i := 0; i < dim-1; i++ {
			fmt.Printf("*")
		}
		//Modulo del medio
	
 		for i := 0; i < dim-1; i++ {
			fmt.Println("*")
		}
		//Ultimo modulo
		
		for i := 0; i < dim; i++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	
}