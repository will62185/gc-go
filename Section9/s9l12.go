// Map - Delete

package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
		"Delete Me":       112211,
	}
	fmt.Println(m)
	fmt.Println("James Bond is", m["James"], "years old.")
	// Prints the entire map using the range function
	for m, v := range m {
		fmt.Println(m, v)
	}
	delete(m, "Delete Me")
	fmt.Println(m)

	if v, ok := m["James"]; ok {
		fmt.Println("value:", v)
		delete(m, "James")
	}
	fmt.Println(m)
}
