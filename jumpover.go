package piscine

func JumpOver(str string) string {
	// Αν το string είναι άδειο -> επιστρέφουμε newline
	if str == "" {
		return "\n"
	}

	result := ""

	for i, r := range str {
		if i%3 == 2 { // κάθε τρίτος χαρακτήρας
			result += string(r)
		}
	}

	// Αν δεν βρέθηκε τρίτος χαρακτήρας -> newline
	if result == "" {
		return "\n"
	}

	return result + "\n"
}
