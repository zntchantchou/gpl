package main

import "fmt"

func main() {
	// var r [3]int = [3]int{1, 2}
	// fmt.Println(r)
	// for _, v := range r {
	// 	fmt.Println("value ", v)
	// }

	r := [...] int{ 99 : -10 }
	for index, v := range r {
		fmt.Println(index, "\t", v)
	}

	nums := [3]int{ 1, 3 }
	fmt.Printf("first element is %x of type %T \n", nums, nums[0])
}