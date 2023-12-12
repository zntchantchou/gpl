package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("goRoutine")
		for i := 0; i < 20; i++ {
			fmt.Println(time.Now().Format("2006-01-02T15:04:05 -07:00:00"))
			fmt.Println(" [LOOP 1]CURRENT ITERATION : ", i, "\n value: ", <-c)
		}
	}()
	time.Sleep(3 * time.Second)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Second function")
		quit <- 0
	}()
	// go func() {
	// 	fmt.Println("goRoutine2")
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(time.Now().String())
	// 		fmt.Println(" [LOOP 2]CURRENT ITERATION : ", i, "\n value: ", <-c)
	// 		// fmt.Println("goRoutine2 ", time.Now().Local().String())
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	fmt.Println("fibonacci")
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
