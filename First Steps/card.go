package main

import "fmt"

type Card struct {
	Rank, Suit string
}

func (c Card) SetSuit(newSuit string) {
	c.Suit = newSuit
}

func (c *Card) SetRank(newRank string) {
	c.Rank = newRank
}

func (c Card) Print() {
	fmt.Println(c.Rank, c.Suit)
}

func PrintFunc(c Card) {
	fmt.Println(c.Rank, c.Suit)
}

func SetRankAlt(c *Card, newRank string) {
	c.Rank = newRank
}

func main() {
	var card = Card{"King", "Spades"}
	card.SetSuit("Diamonds") // Won't change anything
	card.SetRank("Jack")     // Correctly uses pointers
	fmt.Println(card.Rank, card.Suit)

	SetRankAlt(&card, "Ace") // Must be a pointer
	card.Print()
	PrintFunc(card)
}
