package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	// έλεγχος εγκυρότητας
	if !isValidBase(base) {
		for _, r := range "NV" {
			z01.PrintRune(r)
		}
		return
	}

	baseLen := len(base)

	// ειδική περίπτωση για 0
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	// χειρισμός αρνητικών
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	// αποθήκευση ψηφίων
	var digits []rune
	for nbr > 0 {
		remainder := nbr % baseLen
		digits = append([]rune{rune(base[remainder])}, digits...)
		nbr /= baseLen
	}

	// εκτύπωση
	for _, d := range digits {
		z01.PrintRune(d)
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
