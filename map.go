package piscine

func Map(f func(int) bool, a []int) []bool {
	// φτιάχνουμε slice []bool ίδιου μήκους με το a
	result := make([]bool, len(a))

	// γεμίζουμε το result εφαρμόζοντας τη f σε κάθε στοιχείο
	for i, v := range a {
		result[i] = f(v)
	}
	return result
}
