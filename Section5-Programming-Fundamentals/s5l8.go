// Bit Shifting Use: << >>

package main

import (
	"fmt"
)

const (
	_ = iota
	//kb = 1024
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)


	z := x >> 1
	fmt.Printf("%d\t\t%b\n", z, z)

	//kb := 1024
	//mb := 1024 * kb
	//gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\b\n", mb, mb)
	fmt.Printf("%d\t\t%b\b\n", gb, gb)
}
