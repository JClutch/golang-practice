package main

import "testing"

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()

	if len(testDeck) != 16 {
		t.Errorf("Expected deck length of 20 but got %v", len(testDeck))
	}

	if testDeck[0] != "Ace of Spades" {
		t.Errorf("expected the first card in deck to be Ace of Spades but recieved %v", testDeck[0])
	}

	if testDeck[len(testDeck)-1] != "Four of Clubs" {
		t.Errorf("expected the first card in deck to be ace of spades but recieved %v", testDeck[len(testDeck)-1])
	}
}
