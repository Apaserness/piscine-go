package main

import "github.com/01-edu/z01"

func main() {
	var l = 'z'
	for l <='a' {
		z01.PrintRune(l)
		l--
	}
	z01.PrintRune('\n')
}
