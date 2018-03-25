// A Tour of Go Exercise: Loops and Functions
package main

import (
	"fmt"
)

func BaseSqrt(x, z float64) (sqrt float64, iterations int) {
	y := 0.0
	maxIterations := 20
	iterations = 0

	for ; z != y && iterations < maxIterations; iterations++ {
		y = z
		z -= ((z * z) - x) / (2 * z)
		//fmt.Println(z)
	}

	return z, iterations
}

func Sqrt(x float64) {
	fmt.Printf("\nX Value: %v\n", x)
	printData(BaseSqrt(x, 1))
	printData(BaseSqrt(x, x/2))
	printData(BaseSqrt(x, x))
}

func printData(sqrt float64, iterations int) {
	fmt.Printf("Square Root: %.6f \tIterations: %v\n", sqrt, iterations)
}

func main() {
	Sqrt(1)
	Sqrt(2)
	Sqrt(3)
	Sqrt(9)
}
