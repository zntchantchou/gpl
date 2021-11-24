package main

import (
	"fmt"
	"time"
)

func main() {
	go logTimes("ONE", 2)
	logTimes("TWO", 1)
}

// text to log
// delay is the time between two prints
func logTimes(text string, delay int) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(time.Second * time.Duration(delay))
	}
}
