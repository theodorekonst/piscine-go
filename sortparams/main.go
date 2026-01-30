package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func printWord(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]

	// bubble sort
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args)-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// print
	for _, word := range args {
		printWord(word)
	}
}
