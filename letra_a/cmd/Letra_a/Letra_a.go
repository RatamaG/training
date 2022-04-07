package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ingresa un numero ")


		return
	}
	name := os.Args[1]
	if dim, err := strconv.Atoi(name); err == nil {
		
		dimen := make(map[int]string)



	} else{
		fmt.Println("Ingresa un numero, letras no son compatibles ")
	}
} 