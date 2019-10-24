package student

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}

	lice := make([]int, max-min)
	index := 0
	for k := range lice {
		lice[k] = min + index
		index++
	}
	return lice

}
