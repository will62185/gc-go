// Embedded Structs

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	// inner type person gets promted to the outer type secretAgent
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}
	fmt.Println(sa1)
	fmt.Println(p2)
	// access the contents of the struct using dot notation
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.last, p2.age)
}
