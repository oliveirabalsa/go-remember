package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastname string
}

func (p person) getFullName() string {
	return p.name + " " + p.lastname
}

func (p *person) changeFullName(fullName string) {
	chunks := strings.Split(fullName, " ")
	p.name = chunks[0]
	p.lastname = ""
	for index, chunk := range chunks[1:] {
		if index == len(chunks)-2 {
			p.lastname += chunk
		} else {
			p.lastname += chunk + " "
		}
	}

}

func main() {
	newPerson := person{
		name:     "Leonardo",
		lastname: "Oliveira Balsalobre",
	}

	fmt.Println(newPerson.getFullName())
	newPerson.changeFullName("Leonardo Oliveira Balsalobre")
	fmt.Println(newPerson.getFullName())
}
