package main

import (
	"fmt"
)

func main() {
	var sname string
	fmt.Println("What is your name?")
	fmt.Scanln(&sname)
	fmt.Println("Bem Vindo", sname)
}
