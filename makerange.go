package student

import "fmt"

func AppendRange(min, max int) []int {
	if max <= min {
		return nil
	}

	lice = make([]int, max-min)
	index := 0
	for k := range lice {
		lice[k] = min + index
		index++
	}
	return lice

}
func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
