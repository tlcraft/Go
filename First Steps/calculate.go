package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func calculate(x, y int) (int, int) {
	return add(x, y), multiply(x, y)
}

func main() {
	fmt.Println(calculate(42, 13))
	fmt.Println(calculate(2, 5))
}
