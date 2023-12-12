package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne := make(chan string)
	channelTwo := make(chan string)

	go func() {
		time.Sleep(1 * time.Millisecond)
		channelOne <- "Channel one"
	}()

	go func() {
		time.Sleep(2 * time.Millisecond)
		channelTwo <- "Channel two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg2 := <-channelTwo:
			fmt.Println("Msg 2: ", msg2)
		case msg1 := <-channelOne:
			fmt.Println("Msg 1: ", msg1)
		}
	}
}
