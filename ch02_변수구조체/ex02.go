package main

import (
	"fmt"
)

func main() {
	const a = 10

	const (
		b = float32(20.04)
		c = 30
		d = string("40")
	)

	fmt.Printf("%f %d %s\n", b+a, c+a, d+"10")
}
