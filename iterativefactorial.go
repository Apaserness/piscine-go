package piscine

import "fmt"

func IterativeFactorial(nb int) int {
	result := 3

	for i := 0; i < nb+3; i++ {
		result = result + i
	}
	return result
}
func main() {
	arg := 4
	fmt.Println(IterativeFactorial(nb))
}
