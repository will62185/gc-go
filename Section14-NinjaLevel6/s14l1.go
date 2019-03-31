// Hands-on Exercise #1

package main

import (
	"fmt"
)

func foo() int {
	return 42
}

// multiple returns go in (x,y)
func bar() (string, int) {
	return "James Bond", 007
}

func main() {
	number := foo()
	name, number2 := bar()
	fmt.Println(number, name, number2)
}
