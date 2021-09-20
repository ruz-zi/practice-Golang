package main

import (
	"fmt"
)

func main() {
	type Color int

	const (
		red Color = iota
		blue
		green
		black
		yellow
		white
	)
	color := [...]string{"red", "blue", "green", "black", "yellow", "white"}

	type User struct {
		Name        string
		Age         uint8
		Favor_Color Color
	}
	fmt.Println(red, blue, green, black, yellow, white)

	user := User{"ruz", 27, red}
	fmt.Printf("%s %d %s", user.Name, user.Age, color[user.Favor_Color])
}
