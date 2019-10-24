package student

func ConvertBase(nbr, baseFrom, baseTo string) string {
	lenFrom, lenTo := 0, 0
	sliceFrom, sliceTo := []rune(baseFrom), []rune(baseTo)
	for i := range sliceFrom {
		lenFrom = i + 1
	}
	for i := range sliceTo {
		lenTo = i + 1
	}
	n := 0
	r0 := []rune(nbr)
	for _, d := range r0 {
		for i, r := range sliceFrom {
			if d == r {
				n = n*lenFrom + i
			}
		}
	}
	if n == 0 {
		return string(sliceTo[0])
	}
	r := ""
	for n > 0 {
		r = string(sliceTo[n%lenTo]) + r
		n = n / lenTo
	}
	return r
}
