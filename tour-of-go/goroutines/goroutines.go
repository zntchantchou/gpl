package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("runinng in routing")
	say("running right here")
	// s:= "hi";
	// fmt.Println(s);
}
