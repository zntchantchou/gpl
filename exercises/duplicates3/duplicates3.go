package main

import (
	"fmt"
)

func main() {
	// counts := make(map[string]int)
	// fileNames := os.Args[0:]
	// for _, fileName := range fileNames {
	// 	data, err := os.ReadFile(fileName)
	// 	if(err != nil) {
	// 		fmt.Println(err.Error())
	// 	}
	// 	fmt.Println(string(data))
	// 	split := strings.Split(string(data), "\n")
	// 	for elt, 
	// }
 
	var arr []int = []int {1, 2, 3}

	for i, v := range arr {
		fmt.Print(i, v, "\n");
	}

	// size of the array is part of it's type, here declared explicitly 
	arr2 := [3]int {4, 5, 6}


	for i, v := range arr2 {
		fmt.Print(i, v, "\n");
	}

	// here size in the type is determinded at build time by the number of initializers
	arr3 := [...]int {4, 5, 6}

	for i, v := range arr3 {
		fmt.Print(i, v, "\n");
	}
}