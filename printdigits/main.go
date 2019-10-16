package main

import "github.com/01-edu/z01"

func main() {
	var d = 0
	for d <= 9 {
		z01.PrintRune(d)
		d++
	}
	z01.PrintRune('\n')
}
