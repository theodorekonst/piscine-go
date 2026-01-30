package piscine

func StrRev(s string) string {
	runes := []rune(s) // μετατροπή string -> slice από runes
	n := len(runes)

	// swap χαρακτήρων από την αρχή με το τέλος
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes) // slice -> string
}
