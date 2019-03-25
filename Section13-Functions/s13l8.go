// Func Expression

package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	f2 := func(x int) {
		fmt.Println("The year big brother started watching us", x)
	}
	f2(1984)
}
