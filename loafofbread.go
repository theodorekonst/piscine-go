package piscine

func LoafOfBread(str string) string {
	rs := []rune(str)

	// Μέτρησε μη-κενά
	nonSpaces := 0
	for _, r := range rs {
		if r != ' ' {
			nonSpaces++
		}
	}

	// Μόνο κενά ή άδειο: μόνο newline
	if nonSpaces == 0 {
		return "\n"
	}
	// Λιγότερα από 5 μη-κενά: Invalid
	if nonSpaces < 5 {
		return "Invalid Output\n"
	}

	// Χτίσε “φέτες” 5 μη-κενών, μετά κάνε skip τον επόμενο χαρακτήρα (ό,τι κι αν είναι)
	outRunes := []rune{}
	groupCount := 0
	firstWord := true

	for i := 0; i < len(rs); {
		r := rs[i]

		// Αγνόησε spaces όταν μετράς τα 5
		if r == ' ' {
			i++
			continue
		}

		// Αν ξεκινά νέα 5άδα, βάλε κενό ανάμεσα στις “φέτες” (όχι πριν από την πρώτη)
		if groupCount == 0 && !firstWord {
			outRunes = append(outRunes, ' ')
		}
		if groupCount == 0 && firstWord {
			firstWord = false
		}

		outRunes = append(outRunes, r)
		groupCount++
		i++

		if groupCount == 5 {
			// Skip τον επόμενο χαρακτήρα αν υπάρχει (είτε space είτε γράμμα)
			if i < len(rs) {
				i++
			}
			groupCount = 0
		}
	}

	return string(outRunes) + "\n"
}
