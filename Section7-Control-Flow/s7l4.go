// Loop - For Statement
// https://golang.org/ref/spec#For_statements
// init;condition;post
package main

import (
	"fmt"
)

func main() {
	// Similar to a while statement
	x := 0
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}

	/* for x < 10 {
		fmt.Println(x)
		x++
	}*/
	fmt.Println("Loop done.")
}
