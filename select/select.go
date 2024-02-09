package main

import (
	"time"

	"github.com/oliveirabalsa/go-remember/generators"
)

func oMaisRapid(url1, url2, url3 string) string {
	c1 := generators.Titulo(url1)
	c2 := generators.Titulo(url2)
	c3 := generators.Titulo(url3)
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(200 * time.Millisecond):
		return "Todos perderam!"
	}
}

func main() {
	println(oMaisRapid(
		"https://www.cod3r.com.br",
		"https://www.google.com",
		"https://www.mentorcycle.com.br",
	))

}
