// Struct

package main

import (
	"fmt"
)

// A struct is an composite data type. (composite data types, aka, aggregate data types, aka, complex data types).
//Structs allow us to compose together values of different types.

type person struct {
	first string
	last  string
	age   int
}

func main() {
	// created the value of TYPE person assinged to p1
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}
	fmt.Println(p1, p2)
	// access the contents of the struct using dot notation
	fmt.Println(p1.last, p1.age, p2.last, p2.age)
}
