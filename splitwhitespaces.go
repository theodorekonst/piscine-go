package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	word := ""
	for _, r := range s {
		if r != ' ' && r != '\t' && r != '\n' {
			word += string(r)
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}
