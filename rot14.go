package piscine

func Rot14(s string) string {
	var result []rune
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			thesi := r - 'a'
			neathesi := (thesi + 14) % 26
			r = rune('a' + neathesi)
		} else if r >= 'A' && r <= 'Z' {
			thesi := r - 'A'
			neathesi := (thesi + 14) % 26
			r = rune('A' + neathesi)
		}
		result = append(result, r)
	}
	return string(result)
}
