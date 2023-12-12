package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(200 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Print("TICK")
		case <-boom:
			fmt.Print("BOOm !", time.Now().Hour(), ":", time.Now().Minute(), ":", time.Now().Second())
		default:
			fmt.Println(" ..... Default")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
