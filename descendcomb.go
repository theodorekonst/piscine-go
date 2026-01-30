package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := 99; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			// τυπώνουμε τον πρώτο αριθμό
			z01.PrintRune(rune(a/10 + '0'))
			z01.PrintRune(rune(a%10 + '0'))
			z01.PrintRune(' ')
			// τυπώνουμε τον δεύτερο αριθμό
			z01.PrintRune(rune(b/10 + '0'))
			z01.PrintRune(rune(b%10 + '0'))

			// κόμμα + space μόνο αν δεν είναι το τελευταίο
			if !(a == 1 && b == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
