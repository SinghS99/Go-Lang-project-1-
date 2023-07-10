package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if (len(d)) != 16 {
		t.Errorf("Expected deck length of 16 , but got %v", len(d))
	}

	if d[0] != "Heart of ace" {
		t.Errorf("Expected first card Heart of ace, but got %v", d[0])
	}
	if d[len(d)-1] != "joker of three" {
		t.Errorf("Expected last card joker of three but got %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 card but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
