// Array

package main

import (
	"fmt"
)

func main() {
	// Need to specify the size of the array
	var x [5]int
	// Length is part of it's type
	var y [6]int
	fmt.Println(x)
	fmt.Println(y)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	// use slices instead of arrays
}
