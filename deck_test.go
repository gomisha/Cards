package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "4 of Clubs" {
		t.Errorf("Expected last card 4 of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)
	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck got %v", len(loadedDeck))
	}

	os.Remove(fileName)
}
