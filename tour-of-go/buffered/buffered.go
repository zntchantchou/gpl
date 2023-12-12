package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	// this way we avoid overflow and get out when the channel is full
	// sender should be the closer (after sending it is too late!)
	for v := range ch {
		fmt.Print("value ", v)
	}
}
