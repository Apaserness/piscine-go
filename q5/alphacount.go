package student

func AlphaCount(str string) int {
	counter := 0
	len := 0
	for i := range str {
		len = i
	}
	for i := 0; i <= len; i++ {
		counter++
		if !(str[i] >= 'A' && str[i] <= 'Z') && !(str[i] >= 'a' && str[i] <= 'z') {
			counter--
		}
	}
	return counter
}
