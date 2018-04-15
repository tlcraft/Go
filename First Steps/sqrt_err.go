package main

import (
	"fmt"
)

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
