package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i, r := range runes {
		// αν είναι γράμμα ή αριθμός
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if newWord {
				// αν είναι αρχή λέξης → κεφαλαίο
				if r >= 'a' && r <= 'z' {
					runes[i] = r - 32
				} else {
					runes[i] = r
				}
				newWord = false
			} else {
				// αν είναι μέσα στη λέξη → μικρό
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + 32
				} else {
					runes[i] = r
				}
			}
		} else {
			// δεν είναι γράμμα/αριθμός → επόμενος χαρακτήρας θα είναι αρχή λέξης
			newWord = true
		}
	}

	return string(runes)
}
