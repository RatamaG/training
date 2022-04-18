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

	letter := os.Args[1]
	num, err := strconv.Atoi(letter)
	if err != nil {
		fmt.Println("enter a number, letters are not allowed")
	}

	ant1 := 0

	ant2 := 1

	for i := 0; i <= num-1; i++ {

		ant1 += ant2

		ant2 = ant1-ant2

		fmt.Print( ant2 , " . ")
	}

}