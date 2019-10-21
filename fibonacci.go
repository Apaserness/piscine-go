package student

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index > 1 {
		return Fibonacci(index-1) + Fibonacci(index-2)
	} else if index <= 1 {
	return index
	}
	return index
}
