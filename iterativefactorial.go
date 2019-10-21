package main

import "fmt"

func IterativeFactorial(nb int) int {
	if nb >= 1 && nb < 20 {
		result := 1
		for i := 1; i <= nb; i++ {
			result = result * i
		}
		return result
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
}
func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
