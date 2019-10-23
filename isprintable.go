package student

func IsPrintable(str string) bool {
	strlbool := true
	strRune := []rune(str)
	for _, i := range strRune {
		if i >= 0 && i <= 31 {
			strlbool = false
		}
	}
	return strlbool
}
