package main

import "fmt"

type course struct {
	name string
}

func main() {
	var thing interface{}
	fmt.Println(thing)

	thing = 3
	fmt.Println(thing)

	type dynamic interface{}

	var thing2 dynamic = "Opa"
	fmt.Println(thing2)

	thing2 = true
	fmt.Println(thing2)

	thing2 = course{"Golang course"}
	fmt.Println(thing2)
}
