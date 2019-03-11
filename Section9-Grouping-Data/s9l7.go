// Slice - Make

package main

import (
	"fmt"
)

func main() {
	// x length, y capacity
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[9] = 999
	fmt.Println(x)

	// length increases to 11
	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// throws away the old array and creates a new underlying array from 12 to 24
	x = append(x, 3425)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
