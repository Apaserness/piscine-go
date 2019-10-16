package main

import "github.com/01-edu/z01"

func main() {
	d := 0
	for d <= 9 {
		z01.PrintRune(rune(d))
		d++
	}
	z01.PrintRune('\n')
}
