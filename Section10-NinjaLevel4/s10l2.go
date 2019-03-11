// Hands-on Exercise # 2

package main

import (
	"fmt"
)

func main() {
	// Create a composite literal slice of type in with 10 values
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("This is the type of the slice: %T\n", x)
}
