package main

import (
	"fmt"
)

func main() {
	type Employee struct {
		sName     string
		iAge      int
		sFunction string
		salary    float32
	}

	a := Employee{"Paula", 27, "Manager", 10000}
	b := Employee{"Carol", 22, "Analyst", 5500}
	c := Employee{"Peter", 27, "Specialist", 7500}
	fmt.Println("Employee A", a.iAge)
	fmt.Println("Employee B", b.iAge)
	fmt.Println("Employee C", c.iAge)
}
