package main

import (
	"fmt"
)

func main() {
	fmt.Print(1, 2, '3', "\n")
	fmt.Println(1, 2.01, "3")
	fmt.Printf("%d %v %x %b %o", 20, 20, 20, 20, 20)
}
