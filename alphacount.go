package student

import "fmt"

func AlphaCount(str string) int {
	counter := 0
	for index, letter := range str {
		counter++
		fmt.Println(index, letter)

	}
	return counter
}
