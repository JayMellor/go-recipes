package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 5*4 {
		t.Error("Expected 20 cards in deck. Received", len(deck))
	}

	if deck[0] != "Ace of Clubs" {
		t.Error("Expected Ace of Clubs at start of deck. Received", deck[0])
	}

	if deck[len(deck)-1] != "King of Spades" {
		t.Error("Expected King of Spades at end of deck. Received", deck[len(deck)-1])
	}
}

func TestSaveAndLoadDeck(t *testing.T) {
	testFilename := "_decktesting"
	os.Remove(testFilename) // Delete any orphan files

	deck := newDeck()
	error := deck.saveToFile(testFilename)
	if error != nil {
		t.Error("Error saving file:", error)
	}

	loadedDeck, error := newDeckFromFile(testFilename)
	if error != nil {
		t.Error("Error loading file:", error)
	}

	if len(loadedDeck) != 5*4 {
		t.Error("Expected 20 cards in deck. Received", len(deck))
	}

	removeError := os.Remove(testFilename)
	if removeError != nil {
		t.Error("Error deleting file", removeError)
	}
}
