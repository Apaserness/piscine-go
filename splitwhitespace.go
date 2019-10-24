package student

func SplitWhiteSpaces(str string) []string {
	ln := 0
	spaces := false
	for c := range str {
		if spaces && c != 0 && (str[c-1] == '\n' ||
			str[c-1] == '\t' || str[c-1] == ' ') &&
			str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			ln++
		}
		if str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			spaces = true
		}
	}
	ln++

	b := 0
	ans := make([]string, ln)
	ok := true
	for _, c := range str {
		if c == '\n' || c == '\t' || c == ' ' {
			if !ok {
				b++
			}
			ok = true
			continue
		}
		ans[b] = ans[b] + string(c)
		ok = false
	}
	return ans

}
