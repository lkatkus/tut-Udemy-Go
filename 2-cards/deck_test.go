package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected last card to be Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %s", d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	newD, newHand := deal(d, 4)

	if len(newD) != 4 {
		t.Errorf("Expected new deck length of 4, but got %v", len(newD))
	}

	if len(newHand) != 12 {
		t.Errorf("Expected remaing hand length of 12, but got %v", len(newHand))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_deckTesting"

	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove(fileName)
}
