// Loop, Conditional, Modulus

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		// will print out any number divivded by x and has a remander of 0
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
