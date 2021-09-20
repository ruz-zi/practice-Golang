package main

import (
	"fmt"
)

func main() {
	var Hour int
	var Minute int

	fmt.Scanf("%d:%d", &Hour, &Minute)
	fmt.Printf("현재 시각은 %d:%d 입니다.\n", Hour, Minute)
}
