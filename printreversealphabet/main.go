package main

import "github.com/01-edu/z01"

func main() {
	var i = 'z'
	for i >= 'a' {
		z01.PrintRune(i)
		i--
	}
	z01.PrintRune(10)
}
