package student

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	len := 0
	for i := range args {
		len = i
	}
	for i := len; i >= 0; i-- {
		for _, k := range args[i] {
			z01.PrintRune(k)
		}
		z01.PrintRune(10)

	}
}
