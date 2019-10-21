package piscine

import "github.com/01-edu/z01"

func IterativeFactorial(nb int) int {
	result := 0

	for i := 0; i < nb+1; i++ {
		result = result + i
	}
	return result

}
func main() {

	result1 := 0 + 1 + 2 + 3 + 4 + 5 + 6 + 3
	z01.PrintRune(result1)
}
