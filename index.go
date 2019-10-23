package piscine

func Index(s string, toFind string) int {
	runes := []rune(s)
	letfind := []rune(toFind)
	n := rulen(letfind)

	switch {
	case n == 0:
		return 0
	case n == 1:
		return IndexRune(s, letfind[0])
	}

	if rulen(runes) < rulen(letfind) {
		return -1
	}

	for i := range runes {
		if letfind[0] == runes[i] {
			flag := false
			for j := range letfind {
				if runes[i+j] == letfind[j] {
					flag = true
				} else {
					flag = false
				}
			}

			if flag {
				return i
			}
		}
	}

	return -1
}

func IndexRune(s string, r rune) int {
	for i, ch := range []rune(s) {
		if ch == r {
			return i
		}
	}

	return -1
}
