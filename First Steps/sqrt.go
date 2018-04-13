package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Note: a call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
// You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?
// My initial thought: because ErrNegativeSqrt doesn't implement the Stringer interface so an error occurs, causing the Error method to fire, etc.
// Converting the type allows the stringer String method to fire.

func main() {
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-9))
}

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

func Sqrt(x float64) (string, ErrNegativeSqrt) {
	err := ErrNegativeSqrt(x)
	if x < 0 {
		return "", err
	}

	printData(BaseSqrt(x, 1))
	printData(BaseSqrt(x, x/2))
	printData(BaseSqrt(x, x))

	return fmt.Sprintf("\nX Value: %v\n", x), err
}

func printData(sqrt float64, iterations int) {
	fmt.Printf("Square Root: %.6f \tIterations: %v\n", sqrt, iterations)
}
