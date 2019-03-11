// Slice - Underlying Array

/* Three parts to the data structure:
1. Pointer to an underlying array
2. Length of the slice
3. Capacity - length of the array
*/

package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	// y := append(x, 43, 43, 43, 43, 43, 44)// A new underlying array was allocated

	y := append(x[:2], x[5:]...) // the underlying array stores the new values

	fmt.Println(x)
	fmt.Println(y)

}
