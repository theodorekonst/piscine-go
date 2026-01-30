package main

import "github.com/01-edu/z01"

// Ορίζουμε το struct point με δύο πεδία x και y
type point struct {
	x int
	y int
}

// Συνάρτηση που παίρνει pointer σε point και αλλάζει τα πεδία του
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{} // δημιουργεί νέο point με x=0, y=0
	setPoint(points)   // αλλάζει τις τιμές σε 42 και 21
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune('4')
	z01.PrintRune('2')
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune('2')
	z01.PrintRune('1')
	z01.PrintRune('\n')
}
