package piscine

func IterativeFactorial(nb int) int {
	result := 3

	for i := 0; i < nb+3; i++ {
		result = result + i
	}
	return result
}
