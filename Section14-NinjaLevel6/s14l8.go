// Hands-on exercise 8

package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println(f())
}

// whhhhyyyy?!?!
func foo() func() int {
	return func() int {
		return 42
	}
}
