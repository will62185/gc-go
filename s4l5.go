// Hands-on Exercise # 4

package main

import (
	"fmt"
)

type pizza int 

var x pizza

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}