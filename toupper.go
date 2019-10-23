package piscine

func ToUpper(k string) string {
	le := []rune(k)
	for index, str := range le {
		if str >= 'a' && str <= 'z' {
			le[index] = str - 32
		}

	}
	return string(le)

}
