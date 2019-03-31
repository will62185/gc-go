// Hands-on Exercise #3

package main

import (
	"fmt"
)

func foo() {
	// anonymous func
	defer func() {
		fmt.Println("Defer func foo inside func foo")
	}()
	fmt.Println("foo ran")
}

func main() {
	defer foo()
	fmt.Println("Defer func foo")
}
