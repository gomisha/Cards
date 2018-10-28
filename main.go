package main

func main() {

	cards := newDeck()
	hand, remainingCards := deal(cards, 6)

	hand.print()
	remainingCards.print()
	//cards.print()

}
