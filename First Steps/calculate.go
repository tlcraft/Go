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
	printData(calculate(42, 13))
	printData(calculate(2, 5))

	var i, j int = 3, 4
	printData(calculate(i, j))

	k, l := 7, 8.5
	printData(calculate(k, int(l)))

	for a, b := 5, 10; a <= 10; a, b = a+1, b+10 {
		printData(calculate(a, b))
	}
}

func printData(sum, product int) {
	fmt.Printf("Sum: %v \nProduct: %v\n\n", sum, product)
}
