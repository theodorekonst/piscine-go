package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		return
	}

	// αποθηκεύουμε τα ψηφία
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}

	// ταξινομούμε τα ψηφία (bubble sort)
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[j] < digits[i] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}

	// τυπώνουμε
	for _, d := range digits {
		z01.PrintRune(rune(d + '0'))
	}
}
