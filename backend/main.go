package main

import "fmt"

func main() {
	cards := NewDeck()
	hand, remainingCards := deal(cards, 5)
	cards.print()
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
