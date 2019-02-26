// Section 5: Lesson 1 | Numeric Types

package main

import (
	"fmt"
	//"runtime"
)

var x int64
var y float64

func main() {
	x = 42
	y = 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}


/*
// Messing with runtime package
func main(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
*/