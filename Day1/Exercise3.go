package main

import (
	"fmt"
)

func main() {
	var iNumber int
	fmt.Println("Choose one integer number")
	fmt.Scanln(&iNumber)
	if iNumber > 0 {
		fmt.Println("Positivo!", iNumber)
	} else if iNumber < 0 {
		fmt.Println("Negativo!", iNumber)
	} else {
		fmt.Println("Zero!")
	}
}
