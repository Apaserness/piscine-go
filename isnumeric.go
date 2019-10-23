package student

func IsNumeric(str string) bool {
	if str == "" {
		return false
	}
	s := []rune(str)
	for _, num := range s {
		if num >= '0' && num <= '9' {
		} else {
			return false
		}
	}
	return true
}
