package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func calculate(x, y int) (sum, product int) {
	sum = add(x, y)
	product = multiply(x, y)
	return sum, product
}

func main() {
	fmt.Println(calculate(42, 13))
	fmt.Println(calculate(2, 5))
}
