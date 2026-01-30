package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, c := range s { // loop through each character in the string
		z01.PrintRune(c) // print the character
	}
}
