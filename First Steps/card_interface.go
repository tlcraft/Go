package main

import "fmt"

type Card interface {
	GetName() string
}

type StandardCard struct {
	Name string
}

// Implicit Card interface implementation
func (card StandardCard) GetName() string {
	return card.Name
}

func main() {
	var c Card = StandardCard{"Ace of Spades"}
	fmt.Println(c.GetName())
}
