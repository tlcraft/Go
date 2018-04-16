// A Tour of Go Exercise: Errors

package main

import (
	"fmt"
)

// Note: a call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
// You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?
// My initial thought: because ErrNegativeSqrt doesn't implement the Stringer interface so an error occurs, causing the Error method to fire, etc.
// Converting the type allows the stringer String method to fire.

// Answer to the question: https://stackoverflow.com/a/27475316/8094831

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: " + fmt.Sprintf("%.6f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return 0, err
	}

	y := 0.0
	z := 1.0
	maxIterations := 20
	iterations := 0

	for ; z != y && iterations < maxIterations; iterations++ {
		y = z
		z -= ((z * z) - x) / (2 * z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	sqrt, err := Sqrt(3)
	if err == nil {
		fmt.Println(sqrt)
	} else {
		fmt.Println(err)
	}

	sqrt, err = Sqrt(-3)
	if err == nil {
		fmt.Println(sqrt)
	} else {
		fmt.Println(err)
	}
}
