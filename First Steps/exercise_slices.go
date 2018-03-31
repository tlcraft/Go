package main

import (
	"golang.org/x/tour/pic"
)

// Functions as parameters, define a type: https://stackoverflow.com/a/12655719/8094831
type function func(x, y int) uint8

var fn function

func times(x, y int) uint8 {
	return uint8(x * y)
}

func power(x, y int) uint8 {
	return uint8(x ^ y)
}

func parabola(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		pic[i] = make([]uint8, dx)
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			pic[i][j] = fn(i, j)
		}
	}

	return pic
}

//Other solutions: https://gist.github.com/tetsuok/2280162
func main() {
	//fn = times
	fn = power
	//fn = parabola
	pic.Show(Pic)
}
