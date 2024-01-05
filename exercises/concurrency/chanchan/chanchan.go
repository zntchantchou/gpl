package main

import (
	"fmt"
)

func main() {
	chanOne := make(chan int64)
	chanTwo := make(chan int64)

	// time.Sleep(time.Millisecond * 500)
	go func() {
		chanOne <- 1
		close(chanOne)
	}()
	go func() {
		chanTwo <- 2
		close(chanTwo)
	}()
	// chanTwo <- 2
	// time.Sleep(time.Millisecond * 500)
loop:
	for {
		select {
		case v, ok := <-chanOne:
			// fmt.Println("chanOne value ", value)
			fmt.Println("chanOne ", v)
			if !ok {
				fmt.Println("chanOne not Ok")
				break loop
			}

		case v, ok := <-chanTwo:
			// fmt.Println("chanTwo value ", value)
			// fmt.Println("ok ", ok)/
			fmt.Println("chanTwo", v)
			if !ok {
				fmt.Println("chanTwo not Ok")
				break loop
			}
		default:
			fmt.Println("default ")
		}
	}
}
