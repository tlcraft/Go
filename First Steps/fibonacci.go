package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib, num0, num1 := 0, 0, 1
	return func() int {
		if fib == 0 {
			fib++
			return 0
		} else if fib == 1 {
			fib++
			return 1
		}

		num1, num0 = num0+num1, num1

		return num1
	}
}

// More solutions: https://gist.github.com/tetsuok/2281812
// Such as izzlazz's:
func izzlazz_fibonacci() func() int {
	a, b := -1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fibonacci()
	izzlazz := izzlazz_fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Printf("%v %v\n", f(), izzlazz())
	}
}
