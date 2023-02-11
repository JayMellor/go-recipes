package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 5*4 {
		t.Errorf("Expected 20 cards in deck. Received", len(deck))
	}
}
