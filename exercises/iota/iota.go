package main

import "fmt"

func main() {

	// iota is used to create incrementing constants 
	const (
		Red int = iota
		Orange 
		Yellow
		Green 
		Blue 
		Indigo
		Violet
	)

	fmt.Printf("The value of red is %v \n", Red);
	fmt.Printf("The value of Orange is %v \n", Orange);
	fmt.Printf("The value of Indigo is %v \n", Indigo);

	const (
		GBP int = iota
		USD
	)

	symbol := [...]string{ USD:"$", GBP: "£" }
	fmt.Println(symbol[0], USD, GBP)
}
