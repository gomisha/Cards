package main

import (
	"fmt"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4"}
	//	var str strings.Builder

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			// str.WriteString(value)
			// str.WriteString(" of ")
			// str.WriteString(suite)

			//cards = append(cards, str.String)

			cards = append(cards, value+" of "+suite)

		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
