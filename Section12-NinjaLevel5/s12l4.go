// Hands-on Exercise # 4

// Create and use an anonymous struct.

package main

import (
	"fmt"
)

func main() {
	person := struct {
		first string
		last  string
		age   int
	}{

		first: "Rudy",
		last:  "Rudiger",
		age:   25,
	}
	fmt.Println("Hello, my name is", person.first, ",and I'm",person.age,"years old.")

// Example from Todd
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
