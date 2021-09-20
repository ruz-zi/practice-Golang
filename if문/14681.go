package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		fmt.Print(1)
	} else if x < 0 && y > 0 {
		fmt.Print(2)
	} else if x < 0 && y < 0 {
		fmt.Print(3)
	} else {
		fmt.Print(4)
	}
}
