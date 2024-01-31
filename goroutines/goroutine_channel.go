package main

import "time"

func twoTreeFourTimes(base int, c chan int) {
	println(base)
	time.Sleep(time.Second)
	c <- 2 * base
	time.Sleep(time.Second)
	c <- 3 * base
	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go twoTreeFourTimes(2, c)
	println("A")
	go twoTreeFourTimes(3, c)
	a, b := <-c, <-c
	println(a, b)

	println("qualquer coisa")

}
