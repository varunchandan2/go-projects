package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardShapes := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, shape := range cardShapes{   // _ means unused variable
		for _, value := range cardValues{
			cards = append(cards, value+" of "+shape)
		}
	}
	return cards
}

// Create a print function that will print all the cards in the deck
func (d deck) print() {     // d is the receiver
	for i, card := range d { // range helps to iterate through the slice
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {

}
