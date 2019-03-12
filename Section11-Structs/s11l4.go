// Anonymous Structs

package main

import (
	"fmt"
)

func main() {
	// anonymous struct which was person before
	// you can use this if you only need it in a litle area
	p1 := struct {
		first string
		last  string
		age   int
	}{
		// values for struct
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p1)
}
