// Hands-on Exercise # 2

/*
Take the code from the previous exercise, then store the values of type person
in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
*/
package main

import (
	"fmt"
)

type person struct {
	firstname  string
	lastname   string
	favFlavors []string
}

func main() {
	p1 := person{
		firstname: "Rudy",
		lastname:  "Rudiger",
		favFlavors: []string{
			"Mint Cookies and Cream",
			"Strawberry"},
	}
	p2 := person{
		firstname: "Tom",
		lastname:  "Jones",
		favFlavors: []string{
			"Rocky Road",
			"Chocolate",
		},
	}
	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstname)
		fmt.Println(v.lastname)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("-----")

	}

}
