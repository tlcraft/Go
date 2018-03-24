// Write a program that prints the numbers from 1 to 100.
// But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”.
// For numbers which are multiples of both three and five print “FizzBuzz”
// http://wiki.c2.com/?FizzBuzzTest

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	initialFizzBuzz()
	end := time.Now()

	var duration time.Duration = end.Sub(start)

	fmt.Printf("\n** Stats **\nStart: [%v]\nEnd: [%v]\nDuration: [%v]\n", start, end, duration)

	start = time.Now()
	fizzBuzz()
	end = time.Now()

	duration = end.Sub(start)

	fmt.Printf("\n** Stats **\nStart: [%v]\nEnd: [%v]\nDuration: [%v]\n", start, end, duration)
}

func initialFizzBuzz() {
	var fizzBuzz string
	var isNonFizzBuzz bool

	for i := 1; i <= 100; i++ {
		isNonFizzBuzz = true

		if i%3 == 0 {
			fizzBuzz += "Fizz"
			isNonFizzBuzz = false
		}

		if i%5 == 0 {
			fizzBuzz += "Buzz"
			isNonFizzBuzz = false
		}

		if isNonFizzBuzz {
			fizzBuzz += fmt.Sprint(i)
		}

		fizzBuzz += "\n"
	}

	fmt.Print(fizzBuzz)
}

func fizzBuzz() {
	var fizzBuzz string

	for i := 1; i < 101; i++ {
		fizzBuzz = ""

		if i%3 == 0 {
			fizzBuzz += "Fizz"
		}

		if i%5 == 0 {
			fizzBuzz += "Buzz"
		}

		if len(fizzBuzz) == 0 {
			fmt.Println(i)
		} else {
			fmt.Println(fizzBuzz)
		}
	}
}
