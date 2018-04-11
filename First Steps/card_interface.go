package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Card interface {
	GetName() string
	//GetRank() string // Causes an error because the concrete cards don't have a corresponding method - (missing GetRank method)
}

type StandardCard struct {
	Rank string
	Suit string
}

// Implicit Card interface implementation, also I just learnd this playing with concrete types: StandardCard does not implement Card (GetName method has pointer receiver)
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

type UnoCard struct {
	Color  string
	Number int
}

func (uno UnoCard) GetName() string {
	var b bytes.Buffer

	b.WriteString(uno.Color)
	b.WriteString(strconv.Itoa(uno.Number))

	return b.String()
}

func CompareConcreteCardType(i Card) {
	t, ok := i.(UnoCard) // type assertion to grab the interface value's underlying concrete value
	if ok {
		fmt.Println(ok)
		fmt.Printf("%v %T", t, t)
	} else {
		fmt.Println("Not of type UnoCard")
		fmt.Printf("%v %T\n", t, t) // zero values of UnoCard
	}
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

	var uno = UnoCard{"Blue", 3}
	CompareConcreteCardType(i)
	CompareConcreteCardType(uno)
}
