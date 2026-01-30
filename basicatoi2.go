package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		} else {
			// αν υπάρχει χαρακτήρας που δεν είναι ψηφίο → επιστρέφει 0
			return 0
		}
	}
	return result
}
