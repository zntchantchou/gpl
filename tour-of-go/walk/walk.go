package main

import (
	"fmt"
	"slices"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func Walk(t *tree.Tree, c chan int) {
	c <- t.Value
	if t.Left != nil {
		Walk(t.Left, c)
	}
	if t.Right != nil {
		Walk(t.Right, c)
	}
	if t.Left == nil && t.Right == nil {
		fmt.Println("No more nodes ..")
		return
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1values := make(chan int)
	t2values := make(chan int)
	var sequenceOne []int
	var sequenceTwo []int
	go Walk(t1, t1values)
	go Walk(t2, t2values)
	for i := 0; i < 20; i++ {
		select {
		case v := <-t1values:
			sequenceOne = append(sequenceOne, v)
		case v := <-t2values:
			sequenceTwo = append(sequenceTwo, v)
		}
	}
	one := sort.IntSlice(sequenceOne)
	two := sort.IntSlice(sequenceTwo)
	one.Sort()
	two.Sort()
	return slices.Equal(one, two)
}

func main() {
	similarOne := Same(tree.New(1), tree.New(2))
	similarTwo := Same(tree.New(1), tree.New(2))
	fmt.Println("Similar one", similarOne)
	fmt.Println("Similar two", similarTwo)
}
