// Defer
// delays the execution of a function

package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()

}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
