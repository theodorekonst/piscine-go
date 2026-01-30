package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		word := os.Args[i]
		for _, r := range word {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
