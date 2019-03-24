// Unfurling a slice

package main

import (
	"fmt"
)

func main() {
	// xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	// unfurl all of the ints in the slice even though it asks just for ints not []int
	// x := sum(xi...)
	x := sum()
	fmt.Println("The total is", x)

}

// func (r reciver) indentifier(parameter(s)) (returns(s)) {code}

// an unlimited number of ints ...
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in index posistion", i, "We are now adding", v, "to the total which is now", sum)
	}
	return sum
}
