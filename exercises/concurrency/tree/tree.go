package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool

func main() {
	treeChan := make(chan int)
	go Walk(tree.New(5), treeChan)
	for {
		v, ok := <-treeChan
		fmt.Println("ok :", ok)
		fmt.Println("value :", v)
	}
	// close(treeChan)
	// fmt.Println("done")
}

/*
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/
