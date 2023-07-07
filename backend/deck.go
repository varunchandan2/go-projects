package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a struct for card
type Card struct {
	suit  string
	value string
}

// Create a struct for player
type player struct {
	name string
	hand deck
}

// Create a new type of deck which is a slice of strings
type deck []string

func NewDeck() deck {
	cards := deck{}

	cardShapes := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}

	for _, shape := range cardShapes { // _ means unused variable
		for _, value := range cardValues {
			cards = append(cards, value+" of "+shape)
		}
	}
	return cards
}

// Create a print function that will print all the cards in the deck
func (d deck) print() { // d is the receiver
	for i, card := range d { // range helps to iterate through the slice
		fmt.Println(i, card)
	}
}

// this function will convert the slice array type to String type
func (d deck) toString() string {
	return strings.Join([]string(d), ",") // check golang docs for strings
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d { // iterate all the elements through the slice
		newPosition := r.Intn(len(d) - 1) // generates a random number

		d[i], d[newPosition] = d[newPosition], d[i] // swap it out

	}

}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
