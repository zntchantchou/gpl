package main

import "fmt"

func main() {
	nums := []int{1, 4, 5, 3}
	c := make(chan int)
	go sum(nums[len(nums) / 2:], c)
	go sum(nums[:len(nums) / 2], c)

	x, y := <-c, <-c
	fmt.Println("x ", x, " y ", y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
