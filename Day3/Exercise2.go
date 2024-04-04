package main

import (
	"fmt"
)

func main() {
	var message string = "Hello"
	slice := []string{message, "Name1", "Name2", "Name3"}
	secondvalue := slice[2]
	fmt.Println(slice[0], secondvalue)

}
