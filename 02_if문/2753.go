package main

import (
	"fmt"
)

func main() {
	var y int
	fmt.Scan(&y)

	if y%4 == 0 && y%100 != 0 || y%400 == 0 {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
}
