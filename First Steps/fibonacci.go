package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib, num0, num1 := -1, 0, 1
	return func() int {
		fib++
		if fib == 0 {
			return fib
		} else if fib == 1 {
			return fib
		}

		fib = num0 + num1
		num0 = num1
		num1 = fib

		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
