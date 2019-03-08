package main

import "fmt"

var card string

func main() {
	// cards := newDeck()
	cards := newDeckFromFile("testFile1")
	fmt.Println("what is this?", cards)
	// cards.saveToFile("testFile1")
	// hand, remaining := deal(cards, 3)

	// hand.print()
	// remaining.print()
}
