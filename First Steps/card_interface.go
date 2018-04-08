package main

import (
	"bytes"
	"fmt"
)

type Card interface {
	GetName() string
	//GetRank() string // Causes an error because the concrete cards don't have a corresponding method - (missing GetRank method)
}

type StandardCard struct {
	Rank string
	Suit string
}

// Implicit Card interface implementation
func (card *StandardCard) GetName() string {
	if card == nil {
		return "nil"
	} else {
		var b bytes.Buffer

		b.WriteString(card.Rank)
		b.WriteString(" of ")
		b.WriteString(card.Suit)

		return b.String()
	}
}

type TCGCard struct {
	Artist  string
	Title   string
	Defense int
}

func (tcg TCGCard) GetName() string {
	return tcg.Title
}

func main() {
	var c Card = &StandardCard{"Ace", "Spades"}
	fmt.Println(c.GetName())

	var i Card = TCGCard{"Arty", "Goblin", 2}
	fmt.Println(i.GetName())

	var tcg TCGCard = TCGCard{"Art Jr", "Angel", 3}
	fmt.Println(tcg.GetName())

	var iCard Card
	var nilCard *StandardCard
	iCard = nilCard
	fmt.Println(iCard.GetName())
}
