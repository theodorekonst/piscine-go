package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		word := os.Args[i]
		for _, r := range word {
			z01.PrintRune(r)
		}
	}
}
