package main

import (
	"fmt"

	"github.com/varunchandan2/go-projects/practice"
)

// "github.com/varunchandan2/go-projects/cards"

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go practice.CheckLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

/*
	card := cards.NewDeck()
	hand, remainingCards := cards.Deal(card, 5)
	card.Print()
	fmt.Println(hand)
	fmt.Println(remainingCards)

	fmt.Print("\n")
	fmt.Println("Project-2 Check for odd or even number")

	numbers := math.CheckNumber()
	fmt.Println(numbers)

	eb := practice.EnglishBot{}
	sb := practice.SpanishBot{}

	practice.PrintGreeting(eb)
	practice.PrintGreeting(sb)


	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))


	tri := practice.Triangle{Base: 10, Height: 10}
	sq := practice.Square{SideLength: 10}

	practice.PrintArea(tri)
	practice.PrintArea(sq)
*/
