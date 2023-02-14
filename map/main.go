package main

import "fmt"

func main() {

	colours := map[string]string{
		"red":   "#ff0000",
		"blue":  "#4bf745",
		"green": "#432912",
	}

	blankHexes(colours)
	printColours(colours)

}

func blankHexes(colourMap map[string]string) {
	for colour := range colourMap {
		colourMap[colour] = "colour"
	}
}

func printColours(colours map[string]string) {
	for colour, hex := range colours {
		fmt.Println("Hex for", colour, "is", hex)
	}
}
