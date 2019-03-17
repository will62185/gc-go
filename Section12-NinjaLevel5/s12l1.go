// Hands-on Exercise # 1

/* Create your own type “person” which will have an underlying type of “struct”
so that it can store the following data:

first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over
the elements in the slice which stores the favorite flavors.
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
	person1 := person{
		firstname: "Rudy",
		lastname:  "Rudiger",
		favFlavors: []string{
			"Mint Cookies and Cream",
			"Strawberry"},
	}
	person2 := person{
		firstname: "Tom",
		lastname:  "Jones",
		favFlavors: []string{
			"Rocky Road",
			"Chocolate",
		},
	}
	fmt.Println(person1.firstname)
	fmt.Println(person1.lastname)
	for i, v := range person1.favFlavors {
		fmt.Println(i, v)
	}
	fmt.Println(person2.firstname)
	fmt.Println(person2.lastname)
	for i, v := range person2.favFlavors {
		fmt.Println(i, v)
	}

}
