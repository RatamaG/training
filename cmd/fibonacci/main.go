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

	numa := 0

	numb := 1

	cont := 0

	for i := 0; i <= num-1; i++ {

		cont = numa + numb 

		fmt.Print(" ",cont)

		numa = numb
		
		numb = cont
	}

}