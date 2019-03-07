// Hands-on Exercise # 7

package main

import (
	"fmt"
)

func main() {
	x := 43
	if x == 42 {
		fmt.Println(x)
	} else if x <= 42 {
		fmt.Println("My value is: 42")
	} else {
		fmt.Printf("My value is: %v\n", x)
	}

}
