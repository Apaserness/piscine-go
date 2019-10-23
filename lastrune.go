package piscine

func LastRune(s string) rune {
	s1 := []rune(s)
	k := 0
	for _, i := range s1 {
		if i == i {
		}
		k++
	}
	return s1[k-1]
}
