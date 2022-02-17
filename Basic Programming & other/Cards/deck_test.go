package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Length is not 40, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card is not Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Ten of Clubs" {
		t.Errorf("Last card is not Ten of Clubs, got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 40 {
		t.Errorf("Length is not 40, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
