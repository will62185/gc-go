// Slice - Multi-Dimensional Slice

package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	// Slice of a slice of string
	xp := [][]string{jb, mp}
	fmt.Println(xp)

}
