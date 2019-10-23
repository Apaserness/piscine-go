package student

import "os"
import "github.com/01-edu/z01"

func main() {
	name := []rune(os.Args[0])
	for _, r := range name {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
