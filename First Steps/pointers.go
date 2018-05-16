package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Vertices struct {
	list []*Vertex
}

func (v Vertices) Print() {
	for i, v := range v.list {
		fmt.Printf("Index: %v, Vertex: %v\n", i, *v)
	}
}

func (v Vertices) Transform(t int) {
	for _, v := range v.list {
		v.X *= t
		v.Y *= t
	}
}

func main() {
	p := &Vertex{1, 2}      // create a pointer to a Vertex
	fmt.Println("1)\t", *p) // read Vertex through the pointer

	q := &p.X // create a pointer to the X Vertex field
	*q = 8    // set X through the pointer
	p.Y = 10  // set the value of Y

	fmt.Println("2)\t", *p)   // read the Vertex through the pointer
	fmt.Println("3)\t", p)    // read the pointer p
	fmt.Println("4)\t", q)    // read the pointer q
	fmt.Println("5)\t", *q)   // read field X through the pointer q
	fmt.Println("6)\t", &p.X) // read the reference of the field X

	p.X = 15
	fmt.Println("7)\t", p.X)    // read the field from the pointer
	fmt.Println("8)\t", (*p).X) // read the field through the pointer
	fmt.Println("9)\t", *q)     // read X through the pointer
	fmt.Println("10)\t", *p)    // read the Vertex through the pointer

	fmt.Println("  Vertex Array  ")
	vertices.Print()
	vertices.Transform(2)
	vertices.Print()
}

var vertices = Vertices{
	[]*Vertex{
		&Vertex{4, 5},
		&Vertex{10, 3},
		&Vertex{3, 9},
		&Vertex{1, 0},
	},
}
