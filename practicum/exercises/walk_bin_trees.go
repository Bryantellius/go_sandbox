package main

import (
	"fmt"

	"golang.org/x/exp/slices"
	"golang.org/x/tour/tree"
)

// ...From Tour of Go

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	if t.Right != nil {
		Walk(t.Right, ch)
	}

	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	treeVals1 := make([]int, 10)
	treeVals2 := make([]int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		x := <-ch1
		y := <-ch2

		treeVals1[i] = x
		treeVals2[i] = y
	}

	slices.Sort(treeVals1)
	slices.Sort(treeVals2)
	return slices.Compare[[]int](treeVals1, treeVals2) == 0
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
