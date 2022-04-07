package main
import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ingresa un nombre ")

		return
	}
	name := os.Args[1] 
	fmt.Println("",name)
	


} 