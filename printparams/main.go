package student

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	fmt.Println(arguments)
}
