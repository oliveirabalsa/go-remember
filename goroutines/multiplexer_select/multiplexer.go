package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s: falando %d", pessoa, i)
		}
	}()
	return c
}

func juntar(entrada1, entrada2, entrada3 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			case s := <-entrada3:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := juntar(falar("João"), falar("Maria"), falar("Bianca"))
	for i := 0; i < 9; i++ {
		fmt.Println(<-c)
	}
}
