package main

import (
	"fmt"
)

func main() {
	var iNumberOne int
	var iNumberTwo int
	iNumberOne = 7
	fmt.Println("Choose one number")
	fmt.Scanln(&iNumberTwo)
	iSum := iNumberOne + iNumberTwo
	fmt.Println("Sum:", iSum)
}

/*var idade int
idade  20
ou var idade = 20
ou
idade:= 20
ou
var (
	numero1=4
	numero2=8
)
ou
nome, altura, idade = "Rafa", 1.69, 26
*/
