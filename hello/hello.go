package main

import "fmt"

import "log"
import "rsc.io/quote"

import "example/greetings"

func PrintQuote() {
	fmt.Println(quote.Go())
}

func DeferExamples() {
	for i := 0; i < 5; i++ {
		// Run when the function closes
		defer fmt.Printf("%d\n", i)
	}
	PrintQuote()
}

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Leah", "Geoffrey"}

	messages, err := greetings.Helloes(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
