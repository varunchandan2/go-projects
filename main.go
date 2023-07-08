package main

import (
	"fmt"

	"github.com/varunchandan2/go-projects/cards"
	"github.com/varunchandan2/go-projects/math"
)

func main() {
	card := cards.NewDeck()
	hand, remainingCards := cards.Deal(card, 5)
	card.Print()
	fmt.Println(hand)
	fmt.Println(remainingCards)

	fmt.Print("\n")
	fmt.Println("Project-2 Check for odd or even number")

	numbers := math.CheckNumber()
	fmt.Println(numbers)

}
