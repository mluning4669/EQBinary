package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
// Same(tree.New(1),tree.New(1)) should return true
// Same(tree.New(1),tree.New(2)) should return false
// Possibly use a switch statement
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	fmt.Println("hello!")
	t1 := tree.New(1)
	ch := make(chan int)

	go Walk(t1, ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
