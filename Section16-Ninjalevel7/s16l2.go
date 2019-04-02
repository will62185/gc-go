// Hands-on exercise #2

package main

import (
	"fmt"
)

type person struct {
	firstname string
}

func main() {
	p1 := person{
		firstname: "James Bond",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.firstname = "Miss Moneypenny"
	(*p).firstname = "Miss Moneyp"
}
