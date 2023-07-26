package practice

import (
	"fmt"
	"net/http"
)

func CheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep it's up and running!"
}
