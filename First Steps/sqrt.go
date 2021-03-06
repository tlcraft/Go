// A Tour of Go Exercise: Loops and Functions

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	printSqrt(9)
	printSqrt(-2)
	printSqrt(1)
	printSqrt(2)
	printSqrt(3)
	printSqrt(9)
	printSqrt(-9)
}

func printSqrt(x float64) {
	sqrt, err := Sqrt(x)
	if err == nil {
		fmt.Println(sqrt)
	} else {
		fmt.Println(err)
	}
}

func baseSqrt(x, z float64) (sqrt float64, iterations int) {
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

func Sqrt(x float64) (string, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return "", err
	}

	fmt.Println("** SQRT ** ")
	printData(baseSqrt(x, 1))
	printData(baseSqrt(x, x/2))
	printData(baseSqrt(x, x))

	return fmt.Sprintf("\nX Value: %v\n", x), nil
}

func printData(sqrt float64, iterations int) {
	fmt.Printf("Square Root of X: %.6f \tIterations: %v\n", sqrt, iterations)
}
