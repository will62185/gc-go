// Syntax

// organizing your code into little chunks using functions, and packages

// func (r receiver) indentifier(paramters) (returns(s)) {...}

// When you define our func with parameters (if any).
// When call our func and pass in arguments (if any).

package main

import (
	"fmt"
)

func main() {
	foo()
	bar("Rudy")
	s1 := woo("Something Football")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

// func and an identifer "foo"
func foo() {
	fmt.Println("Hello from foo")
}

// everything in GO is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello from bar,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

// string, bool are the returns in parans because more than one
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}
