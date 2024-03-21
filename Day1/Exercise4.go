package main

import (
	"fmt"
)

func main() {
	var iCont int
	for iCont = 10; iCont != 0; iCont-- {
		fmt.Println("Contagem regressiva", iCont)
	}
	fmt.Println("Arrasou!")
}
