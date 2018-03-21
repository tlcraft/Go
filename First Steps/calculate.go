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

	var i, j int = 3, 4
	fmt.Println(calculate(i, j))

	k, l := 7, 8.5
	fmt.Println(calculate(k, int(l)))
}
