package main

import (
	"fmt"
)

func main() {
	var h, m int
	fmt.Scan(&h, &m)
	m += h*60 - 45
	if m < 0 {
		m += 1440
	}
	fmt.Print(m/60, m%60)
}
