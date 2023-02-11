package main

import "fmt"
import "os"
import "strings"
import "math/rand"
import "time"

type deck []string

func newDeck() deck {
	suits := [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}
	values := [5]string{"Ace", "Two", "Jack", "Queen", "King"}
	newDeck := deck{}

	for _, suit := range suits {
		for _, value := range values {
			newDeck = append(newDeck, value+" of "+suit)
		}
	}

	return newDeck
}

func (dk deck) print() {
	for _, card := range dk {
		fmt.Println(card)
	}
}

func (dk deck) myShuffle() {
	now := time.Now().UnixNano()
	source := rand.NewSource(now)
	rand := rand.New(source)
	for index := range dk {
		newPosition := rand.Intn(len(dk) - 1)
		dk[index], dk[newPosition] = dk[newPosition], dk[index]
	}
}

func (dk deck) shuffle() {
	now := time.Now().UnixNano()
	source := rand.NewSource(now)
	rand := rand.New(source)
	rand.Shuffle(len(dk), func(fst, scd int) {
		dk[fst], dk[scd] = dk[scd], dk[fst]
	})
}

// Multiple return values
func deal(dk deck, handSize int) (deck, deck) {
	return dk[:handSize], dk[handSize:]
}

func (dk deck) saveToFile(filename string) error {
	deckString := strings.Join(dk, ",")
	return os.WriteFile(filename, []byte(deckString), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	fileBytes, error := os.ReadFile(filename)

	if error != nil {
		return deck{}, error
	}

	deckStrings := strings.Split(string(fileBytes), ",")
	return deck(deckStrings), error
}
