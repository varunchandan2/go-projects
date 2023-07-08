package math

import "fmt"

type data []int

func CheckNumber() data {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, d := range data {

		if d%2 == 0 {
			fmt.Println("The number is even")

		} else {
			fmt.Println("The number is odd")
		}

	}
	return data
}
