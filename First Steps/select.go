package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func twos(c, done chan int) {
	z := 0
	for {
		select {
		case c <- z:
			z += 2
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	ch := make(chan int)
	done := make(chan int)

	go func() {
		for j := 0; j < 20; j++ {
			fmt.Println(<-ch, j)
		}
		done <- 0
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c, i)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
	twos(ch, done)

	fmt.Println("end")
}
