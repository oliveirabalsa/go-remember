package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}

	}
	return true
}

func primos(n int, ch chan int) {
	start := 2
	for i := 0; i < n; i++ {
		for primo := start; ; primo++ {
			if isPrimo(primo) {
				ch <- primo
				start = primo + 1
				time.Sleep(time.Microsecond * 180)
				break
			}
		}
	}
	close(ch)
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c)
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("End!")
}
