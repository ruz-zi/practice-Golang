package main

import (
	"fmt"
)

func main() {
	a := 20
	fmt.Printf("%T %d\n", a, a)
	fmt.Printf("%T %f\n", 20.04, 20.04)
	fmt.Printf("%t\n", a == 20)
}
