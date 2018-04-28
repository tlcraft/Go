package main

import (
	"fmt"
	"time"
)

func print(s string) {
	fmt.Println(s, time.Now())
}

func add(x, y int, c chan int) {
	time.Sleep(1000 * time.Millisecond)
	c <- x + y
}

func add2(x, y int, c chan int) {
	defer say("add2")
	c <- x + y
	close(c)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func main() {
	defer say("Good-bye")

	c := make(chan int)
	deferCh := make(chan int)

	go add(1, 2, c)
	go add2(2, 3, deferCh)

	print("One")

	go say("world")
	go say("hello")

	go print("Four")
	defer print("Two")
	go print("Three")

	sum := <-c
	sumCh := <-deferCh

	fmt.Println(sum)
	fmt.Println(sumCh)

	fmt.Println("End")
}
