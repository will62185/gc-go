// Hands-on Exercise # 4

package main

import (
	"fmt"
)

func main() {
	birthday := 1985
	for {
		if birthday > 2019 {
			break
		}
		fmt.Println(birthday)
		birthday++
	}

}
