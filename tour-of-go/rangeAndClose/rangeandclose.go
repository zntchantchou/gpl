package main

import "fmt"

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for v := range c {
		fmt.Print(v)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		// send 0 to channel
		x, y = y, x+y
	}
	close(c)
}
