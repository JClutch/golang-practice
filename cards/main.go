package main

import (
	"fmt"
)

var card string

func main() {
	cards := newDeck()
	fmt.Println(cards)
}
