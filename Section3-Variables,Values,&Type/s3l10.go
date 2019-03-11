// Conversion, Not Casting

package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {

// TYPE is static it can't change

	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Convert aka Casting: type hotdog b to type int a
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
