package main

import (
	"fmt"
)

func main() {
	var sname string
	var nage int
	nage = 27
	fmt.Println("What is your name?")
	fmt.Scanln(&sname)
	fmt.Println("Welcome", sname, "tem", nage)
}
