// Slice - Deleting From a Slice

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

	// Delete 7, 8 out of slice x
	/* give me everything but not including 2 then append evertyhing
	from posistion 4 and on.*/

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
