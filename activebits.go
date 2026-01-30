package piscine

func ActiveBits(n int) int {
	count := 0
	for n > 0 {
		// Ελέγχει αν το τελευταίο bit είναι 1
		if n%2 == 1 {
			count++
		}
		// Μετακινεί τον αριθμό δεξιά (διαιρεί με 2)
		n = n / 2
	}
	return count
}
