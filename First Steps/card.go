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

func SetRankFunc(c *Card, newRank string) {
	c.Rank = newRank
}

func (c Card) Print() {
	fmt.Println(c.Rank, c.Suit)
}

func PrintFunc(c Card) {
	fmt.Println(c.Rank, c.Suit)
}

func main() {
	fmt.Println("Create two cards")
	card, p := Card{"King", "Spades"}, &Card{"3", "Hearts"}

	card.Print()
	p.Print() // Methods can use either a pointer or a value, pointers are interpreted in this case by Go as values (*p)
	(*p).Print()
	//*p.Print() // Error

	fmt.Println("Alter the first card")
	card.SetSuit("Diamonds") // Won't change anything, this method doesn't use pointers
	card.SetRank("Jack")     // Go interprets the pointer &card for me
	//&card.SetRank("Queen") // Error
	(&card).SetRank("Queen")
	card.Print()
	PrintFunc(card)

	fmt.Println("Alter the second card")
	p.SetSuit("Clubs") // Still won't change anything
	p.SetRank("10")    // Correctly uses pointers
	p.Print()
	PrintFunc(*p) // Needs the underlying value of the pointer

	fmt.Println("Use a function to edit the rank")
	SetRankFunc(&card, "Ace") // Must pass in a pointer
	card.Print()
	PrintFunc(card)

	SetRankFunc(p, "5")
	p.Print()
	PrintFunc(*p) // Needs the underlying value of the pointer
}
