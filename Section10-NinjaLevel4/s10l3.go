// Hands-on Exercise # 3

package main

import (
	"fmt"
)

func main() {
	// Create a composite literal slice of type in with 10 values
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// slicing on the slice
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)
}
