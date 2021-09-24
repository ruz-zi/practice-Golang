package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	for ; n > 0; n-- {
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
