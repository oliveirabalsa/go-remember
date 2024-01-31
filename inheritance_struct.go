package main

import "fmt"

type car struct {
	name         string
	currentSpeed int
}
type ferrari struct {
	car
	turboOn bool
}

func init() {
	f := ferrari{}
	f.name = "F40"
	f.currentSpeed = 0
	f.turboOn = true

	fmt.Println("A ferrari", f.name, "está com o turbo ligado?", f.turboOn)
}
