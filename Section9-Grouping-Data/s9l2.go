// Slice - Composite Literal

package main

import (
	"fmt"
)

func main() {
	// COMPOSITE literal: x := type{values}
	x := []int{4, 5, 6, 7, 8, 42} // slice of int
	// a SLICE allows you to group together VALUES of the same TYPE
	fmt.Println(x)
}
