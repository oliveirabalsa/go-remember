package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // write

	<-ch // read

	ch <- 2
	fmt.Println(<-ch)

}
