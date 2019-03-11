// Map - Add Element & Range

package main

import (
	"fmt"
)

func main() {
	// map [key=string]value=string
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println("James Bond is", m["James"], "years old.")

	// Returns 0 because there is no entry in the map for odd job
	fmt.Println(m["Odd Job"])

	// Check for the value and see if it exists in the map
	// comma ok idiom
	/*v, ok := m["Barnabas"]
	fmt.Println(v, ok)*/

	// Add key Todd value 33 to map
	m["todd"] = 33

	/* Does not print since it doesn't exist
	if v, ok := m["Barnabas"]; ok {
		fmt.Println("This is the IF Print", v)
	}
	// This does run
	if v, ok := m["James"]; ok {
		fmt.Println("This is the IF Print", v)
	}*/
	// Print the key and value of the entire map using the range function
	for k, v := range m {
		fmt.Println(k, v)
	}
	xi := []int{4, 5, 7, 8, 9, 42}
	// range of slice xi and print it all out
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
