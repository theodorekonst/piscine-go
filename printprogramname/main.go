package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	progPath := os.Args[0]
	lastSlash := -1

	// βρίσκουμε τη θέση του τελευταίου '/'
	for i := 0; i < len(progPath); i++ {
		if progPath[i] == '/' {
			lastSlash = i
		}
	}

	// εκτύπωση χαρακτήρα–χαρακτήρα από το όνομα
	for _, r := range progPath[lastSlash+1:] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
