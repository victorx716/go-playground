package main

import (
	"os"
	"testing"
)

// t is test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but received %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the 1st card to be Ace of Spades, received %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, instead got %v", d[len(d)])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
