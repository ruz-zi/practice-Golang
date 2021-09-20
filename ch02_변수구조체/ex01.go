package main

import (
	"fmt"
)

func main() {
	var a int
	var b uint8 = 10
	var c = uint8(20)
	var d = b + c
	e := uint(40)
	var f float32 = 50
	var g = float64(60)
	var h string = " 70 "
	var (
		i int = 80
		j int = 90
	)
	fmt.Print(a, b, c, d, e, f, g, h, i, j)
}
