package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)

	switch {
	case a >= 90:
		fmt.Print("A")
	case a >= 80:
		fmt.Print("B")
	case a >= 70:
		fmt.Print("C")
	case a >= 60:
		fmt.Print("D")
	default:
		fmt.Print("F")
	}
}
