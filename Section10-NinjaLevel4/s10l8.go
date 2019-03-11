// Hands-on Exercise # 8

package main

import (
	"fmt"
)

/* Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things.
Store three records in your map. Print out all of the values, along with their index position in the slice.

		`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`

		`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`

		`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

*/

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	// fmt.Println(m)
	// loop through the key,and values for the map using range & print out the key
	for key, value := range m {
		fmt.Println("This is the record for", key)
		for index, value2 := range value { // loop through the index[] value2 using range
			fmt.Println("\t", index, value2)
		}
	}
}
