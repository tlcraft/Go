package main

import (
	"fmt"
	"image"
	"image/color"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	m.Set(0, 0, color.RGBA{0, 100, 255, 10})
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	fmt.Println(m)
	fmt.Println(m.ColorModel())
	fmt.Println(m.At(0, 0))
	fmt.Println(m.At(10, 10))
}
