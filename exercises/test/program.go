package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	num := os.Args[2]
	if num != "" {
		numAsString, err := strconv.ParseInt(num, 10, 8)
		if err != nil {
			fmt.Println(err)
			return
		}
		repeat(os.Args[1], int(numAsString))
	}
}

func repeat(s string, num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(s)
	}
}
