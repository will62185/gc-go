// Hands-on Exercise # 1

package main

import (
	"fmt"
)

func main() {
	// Create a composite literal array holding 5 values
	x := [5]int{0, 1, 2, 3, 4}
	for i, v := range x {
		fmt.Println(i, v)
	}
	// Print out the type of the array
	fmt.Printf("%T\n", x)
}
