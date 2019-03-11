// Iota

package main

import (
	"fmt"
)

// Either way is fine to show iota
const (
	a = iota
	b
	c
	d
	e
	f
)

// New const block starts the increment over from 0
const (
	g = iota
	h = iota
	i = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

}