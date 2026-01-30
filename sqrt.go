package piscine

func Sqrt(nb int) int {
	for r := 1; r <= nb; r++ {
		if r*r == nb {
			return r
		}
	}
	return 0
}
