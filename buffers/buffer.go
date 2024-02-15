package main

import "fmt"

func routine(ch chan int) {
	fmt.Println("hello")
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
}

func main() {
	channel := make(chan int, 3)

	go routine(channel)

	fmt.Println(<-channel)
}
