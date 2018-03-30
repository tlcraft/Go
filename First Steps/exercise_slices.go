package main

import (
	"math"

	"golang.org/x/tour/pic"
)

func times(x, y int) uint8 {
	return uint8(x * y)
}

func power(x, y int) uint8 {
	return uint8(int(math.Pow(float64(x), float64(y))))
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
			pic[i][j] = parabola(j, i)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
