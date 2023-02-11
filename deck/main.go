package main

import "fmt"
import "os"

func main() {
	cards := newDeck()
	error := cards.saveToFile("cards.cd")
	if error != nil {
		fmt.Println("Error saving file:", error)
		os.Exit(1)
	}

	deck, error := newDeckFromFile("cards.cd")
	if error != nil {
		fmt.Println("Error reading file:", error)
		os.Exit(1)
	}

	deck.myShuffle()
	deck.print()
}
