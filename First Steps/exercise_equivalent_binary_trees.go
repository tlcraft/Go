package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	RecursiveWalk(t, ch)
	close(ch)
}

// Great point: https://stackoverflow.com/a/12224111/8094831
func RecursiveWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		ch <- t.Value

		if t.Left != nil {
			RecursiveWalk(t.Left, ch)
		}

		if t.Right != nil {
			RecursiveWalk(t.Right, ch)
		}
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	m := make(map[int]int)
	var areSame bool = true

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok := <-ch1
		if ok {
			m[v1] += 1
		} else {
			break
		}
	}

	for {
		v2, ok := <-ch2
		if ok {
			m[v2] += 1
		} else {
			break
		}
	}

	for _, v := range m {
		if v != 2 {
			areSame = false
			break
		}
	}

	return areSame
}

func main() {
	ch := make(chan int)
	tree0 := tree.New(1)
	go Walk(tree0, ch)
	fmt.Println(tree0)

	for {
		u, ok := <-ch
		if ok {
			fmt.Println(u)
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
