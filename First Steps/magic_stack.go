package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		switch {
		case i == 0:
			defer fmt.Printf("Upkeep %v\n", i)
		case i == 9:
			defer fmt.Printf("End turn %v\n", i)

		case i%6 == 0:
			defer fmt.Printf("Declare Blockers %v\n", i)
		case i%5 == 0:
			defer fmt.Printf("Declare Attackers %v\n", i)
		case i%4 == 0:
			defer fmt.Printf("Play creature %v\n", i)
		case i%3 == 0:
			defer fmt.Printf("Play instant %v\n", i)
		case i%2 == 0:
			defer fmt.Printf("Tap for Mana %v\n", i)

		default:
			defer fmt.Printf("Thinking... %v\n", i)
		}
	}

	fmt.Println("Magic The Gathering - Switch and Defer Stack")
}
