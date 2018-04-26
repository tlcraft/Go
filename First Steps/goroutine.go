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
	time.Sleep(1000 * time.Millisecond)
	c <- x + y
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func main() {
	c := make(chan int)
	ch := make(chan int)
	go add(1, 2, c)
	defer add(2, 3, ch)
	print("One")
	defer say("Good-bye")

	go say("world")
	go say("hello")

	go print("Four")
	defer print("Two")
	go print("Three")
	sum := <-c
	sumCh := <-ch // blows up I presume because ch won't be available until after this surrounding function returns according to defer
	fmt.Println(sum)
	fmt.Println(sumCh)

	fmt.Println("End")
}
