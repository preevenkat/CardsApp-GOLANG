package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length is  %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is  %v", d[0])
	}

	if d[len(d)-1] != "Four of Hearts" {
		t.Errorf("Expected last card is  %v", d[len(d)-1])
	}

}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length is  %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
