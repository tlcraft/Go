package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := &Vertex{1, 2}
	fmt.Println(*p)

	q := &p.X
	*q = 8
	p.Y = 10

	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(q)
}
