package main

import (
	"bytes"
	"fmt"
)

type Card interface {
	GetName() string
}

type StandardCard struct {
	Rank string
	Suit string
}

// Implicit Card interface implementation
func (card StandardCard) GetName() string {
	var b bytes.Buffer

	b.WriteString(card.Rank)
	b.WriteString(" of ")
	b.WriteString(card.Suit)

	return b.String()
}

func main() {
	var c Card = StandardCard{"Ace", "Spades"}
	fmt.Println(c.GetName())
}
