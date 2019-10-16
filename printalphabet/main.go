package main

import "github.com/01-edu/z01"

func main() {
	
	var l = 'a'
	for l <= 'z' {
		z01.PrintRune(l)
		l++
	}
	z01.PrintRune('\n')
}
