package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var x int
	var input string
	var err error

	fmt.Println("Please enter in a seed value.")
	fmt.Scanln(&input)

	x, err = strconv.Atoi(input)
	if err != nil {
		// handle error
		fmt.Println("Please enter in a number next time. Exception:", err)
	} else {
		fmt.Println("My favorite number is", rand.Intn(x))
	}
}
