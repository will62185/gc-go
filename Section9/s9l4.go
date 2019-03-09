// Slice - Slicing a Slice

package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1]) // access by index
	fmt.Println(x)
	fmt.Println(x[1:]) // : Slicing the Slice
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

}
