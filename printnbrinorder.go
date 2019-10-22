package student

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		k = append(k, n) //prisvoit
	}
	index := 0
	for n > 0 {
		k = append(k, n%10)
		n = n / 10
		index++
	}
	for i := 0; i < index; i++ {
		for j := i + 1; j < index; j++ {
			if k[i] > k[j] {
				t := k[i]
				k[i] = k[j]
				k[j] = t

			}
		}
	}
	for i := range k {
		z01.PrintRune(rune(48 + k[i]))
	}
