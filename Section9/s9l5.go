// Slice - Append to a slice

package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...) // ... variadic parameter - unlimited number of this type
	fmt.Println(x)
	// Messing around below
	y = append(y, x...)
	fmt.Println(y)
}
