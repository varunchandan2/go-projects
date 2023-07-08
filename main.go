package main

import (
	"fmt"

	"github.com/varunchandan2/go-projects/cards"
)

func main() {
	card := cards.NewDeck()
	hand, remainingCards := cards.Deal(card, 5)
	card.Print()
	fmt.Println(hand)
	fmt.Println(remainingCards)

}

//Extra Notes

/*
	// Type conversion for any data
	greeting := "Hello World!"
	fmt.Println([]byte(greeting))
*/

/*
Slice arrays all should have the same data type and can grow and shrink
Arrays defined has defined length, and you have to declare it everytime you use

*/
