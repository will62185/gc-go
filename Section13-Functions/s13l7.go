// Anonymous Func

package main

import (
	"fmt"
)

func main() {
	foo()
	// anonymous funcs
	func() {
		fmt.Println("Anonymous func ran")
	}() // no argument at the end of the func
	func(x int) {
		fmt.Println("The meaning of life", x)
	}(42) // argument 42
}

func foo() {
	fmt.Println("foo ran")
}
