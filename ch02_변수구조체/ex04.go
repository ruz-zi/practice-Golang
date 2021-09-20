package main

import (
	"fmt"
)

const (
	Man   bool = false
	Woman bool = true
)

func main() {
	type Human struct {
		Name   string
		Age    uint8
		Height uint8
		Sex    bool
	}
	type Student struct {
		Humaninfo Human
		Grade     uint8
		School    string
	}
	type Teacher struct {
		Human
		School string
	}

	a := Student{Human{"ruz", 27, 180, Man}, 4, "Inha"}
	b := Teacher{Human{"kim", 30, 167, Woman}, "Inha"}
	fmt.Print(a.Humaninfo.Name, a.School)
	fmt.Print(b.Age, b.Height)
}
