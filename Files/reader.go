package files

import "fmt"

func ReadAFile() {
	file := make([]int, 0)

	file = append(file, 22)
	fmt.Println(file)
}
