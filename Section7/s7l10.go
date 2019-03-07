// Conditional - Switch Statement

package main

import (
	"fmt"
)

func main() {
	name := "Bond"
	switch name /*"Bond"*/ {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("Miss Money, or Bond, or Dr No.")
	case "M":
		fmt.Println("M.")
	case "Q":
		fmt.Println("Q.")
	default:
		fmt.Println("this is default")
		/*case (2 == 4):
			fmt.Println("This should not print 002.")
		case (3 == 3):
			fmt.Println("This should print.")
			fallthrough
		case (4 == 4):
			fmt.Println("Also true, does it print?")
			fallthrough
		case (7 == 9):
			fmt.Println("Not tue 1")
			fallthrough
		case (11 == 14):
			fmt.Println("Not true 2.")
			fallthrough
		case (15 == 15):
			fmt.Println("Not true 15.")
			fallthrough
		default:
			fmt.Println("this is default")*/
	}
}
