package main

import "fmt"

func IterativeFactorial(nb int) int {
	result := 0

	for i := 0; i < nb+1; i++ {
		result = result + i
	}
	return result

}
func main() {

	result1 := 0 + 1 + 2 + 3 + 4 + 5 + 6 + 3
	fmt.Println(result1)
}
