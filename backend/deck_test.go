package main

import "testing"

// Check for x number of cards
// Check the first card is the same
// Check for the last card

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	if len(d) != 52 {
		t.Errorf("Expected length 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card as Ace of spade but got %v", d[0])
	}

	if d[len(d)-1] != "Jack of Hearts" {
		t.Errorf("Expected the last card Jack of Hearts but got %v", d[len(d)-1])
	}

}
