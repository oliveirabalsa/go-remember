package main

import (
	"fmt"
	"github.com/oliveirabalsa/go-remember/generators"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	res := juntar(
		generators.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		generators.Titulo("https://www.mentorcycle.com.br", "https://www.youtube.com"),
	)

	fmt.Println(<-res, "|", <-res)
	fmt.Println(<-res, "|", <-res)
}
