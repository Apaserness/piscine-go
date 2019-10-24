package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	for _, r := range os.Args[1:] {
		PrintRuneArray([]rune(r))
		z01.PrintRune(10)
	}

}
func PrintRuneArray(runes []rune) {
	for _, ch := range runes {
		_ = z01.PrintRune(ch)
	}
}
