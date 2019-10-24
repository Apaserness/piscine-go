package student

import "fmt"

func AppendRange(min, max int) []int {
	var slice []int
	for i := min; i < max; i++ {

		slice = append(slice, i)
	}
	if max <= min {
		return nil
	}
	return slice

}
