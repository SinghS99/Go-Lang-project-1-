package main

import "testing"

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
