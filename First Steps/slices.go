package main

import "fmt"

func main() {
	s := make([]int, 6)
	r := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	for i := range r {
		s = append(s, r[i])
	}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:6]
	printSlice(s)

	// Drop its first five values.
	t := s[5:12]
	printSlice(s)
	printSlice(t)
	printSlice(r)

	s = s[0:12]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
