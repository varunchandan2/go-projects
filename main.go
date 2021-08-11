package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of spades")

	for i, card := range cards { // range says to iterate through the slice
		fmt.Println(i, card)
	}



}

func newCard() string {
	return "Five of Diamonds"
}





/*
Slice arrays all should have the same data type and can grow and shrink
Arrays defined has defined length, and you have to declare it everytime you use

 */