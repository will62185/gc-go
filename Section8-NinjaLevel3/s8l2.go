// Hands-on Exercise # 2

package main

import (
	"fmt"
)

func main() {
	// for init; conditionl post {}
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
