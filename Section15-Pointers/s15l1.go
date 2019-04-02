// What Are Pointers?

/*
All values are stored in memory.
Every location in memory has an address.
A pointer is a memory address.
*/

package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	// & is an operator show the address in memory where a is stored
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	// pointer to an int 'pointer to a type"
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an address when you have the address
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)

}
