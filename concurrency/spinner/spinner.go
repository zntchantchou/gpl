package main

import (
	"fmt"
	"time"
)

// main is our main goroutine, with one child go routine
// once main goroutine is done, the child goroutine (call to spinner) gets killed
func main() {
	go spinner(100 * time.Millisecond)
	const n = 60
	fib := fibonacci(n)
	fmt.Printf("Fibonacci of %d is %d", n, fib)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fibonacci(num int) int {
	if num < 2 {
		return num
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
