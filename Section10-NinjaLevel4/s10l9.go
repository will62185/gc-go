// Hands-on Exercise # 9

package main

import (
	"fmt"
)

//Using the code from the previous example, add a record to your map.
// Now print the map out using the “range” loop

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	// Add job_odd to the map
	m["job_odd"] = []string{"Hats", "Suits", "Candy Canes"}

	// fmt.Println(m)
	// loop through the key,and values for the map using range & print out the key
	for key, value := range m {
		fmt.Println("This is the record for", key)
		for index, value2 := range value { // loop through the index[] value2 using range
			fmt.Println("\t", index, value2)
		}
	}
}
