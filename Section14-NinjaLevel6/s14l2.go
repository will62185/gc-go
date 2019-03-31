// Hands-on Exercise #2

package main

import (
	"fmt"
)

//variadic parameter
func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := foo(ii...)
	fmt.Println(n)

	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n2 := bar(ii2)
	fmt.Println(n2)
}
