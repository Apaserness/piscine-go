package student

func ToUpper(s string) string {
	srunes := []rune(s)
	var res rune
	for i := 0; i < len(s); i++ {
		if srunes[i] >= 'a' && srunes[i] <= 'z' {
			res = srunes(s[i] - 32)
			srunes[i] = res
		}
	}
	return string(srunes)
}
