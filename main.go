package main

import (
	"fmt"
)

func main() {

	cards := newDeck()
	// hand, remainingCards := deal(cards, 6)

	// hand.print()
	// remainingCards.print()

	// greeting := "Hi there"
	// fmt.Println([]byte(greeting))

	fmt.Println(cards.toString())

	cards.saveToFile("cards.txt")
}
