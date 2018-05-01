package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		ch <- t.Value

		if t.Left != nil {
			Walk(t.Left, ch)
		}

		if t.Right != nil {
			Walk(t.Right, ch)
		}
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	// TODO try using a capacity and range instead
	for i := 0; i < 10; i++ {
		i, ok := <-ch
		if ok {
			fmt.Println(i)
		} else {
			break
		}
	}

	SameFuncTest(tree.New(1), tree.New(1))
	SameFuncTest(tree.New(1), tree.New(2))
	SameFuncTest(tree.New(3), tree.New(3))
	SameFuncTest(tree.New(4), tree.New(3))
}

func SameFuncTest(tree1, tree2 *tree.Tree) {
	areSame := Same(tree1, tree2)
	fmt.Println("\ntree1 and tree2 are the same:", areSame)
	fmt.Println("tree1", tree1.String())
	fmt.Println("tree2", tree2.String())
}
