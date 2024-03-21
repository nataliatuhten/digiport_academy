package main

import (
	"fmt"
)

func main() {
	var iNumber1 int
	var iNumber2 int
	var iOperator string

	fmt.Println("Choose one integer positive number")
	fmt.Scanln(&iNumber1)
	fmt.Println("Choose one more integer positive number")
	fmt.Scanln(&iNumber2)
	fmt.Println("Choose one operation", "+ - * /")
	fmt.Scanln(&iOperator)
	if iNumber1 > 0 && iNumber2 > 0 {
		if iOperator == ("+") {
			iSum := iNumber1 + iNumber2
			fmt.Println(iSum)
		} else if iOperator == "-" {
			iSum := iNumber1 - iNumber2
			fmt.Println(iSum)
		} else if iOperator == "*" {
			iSum := iNumber1 * iNumber2
			fmt.Println(iSum)
		} else if iOperator == "/" {
			iSum := iNumber1 / iNumber2
			fmt.Println(iSum)
		} else {
			fmt.Println("Choose a valid operatior!")
		}
	} else {
		fmt.Println("Choose a positive number!")
	}
}
