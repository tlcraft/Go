package main

import (
	"fmt"
)

type MyInt int

func (i MyInt) double() int {
	return int(i) + int(i)
}

func (i MyInt) add(x int) int {
	return int(i) + x
}

// We need a pointer to the type: https://stackoverflow.com/a/8021805/8094831
func (i *MyInt) set(x int) {
	*i = MyInt(x)
}

func (i *MyInt) setMyInt(x MyInt) {
	*i = x
}

func main() {
	var x MyInt = 5
	fmt.Println(x, x.double())
	fmt.Println(x, x.add(10))
	x.set(12)
	fmt.Println(x)
	x.setMyInt(2)
	fmt.Println(x)
}
