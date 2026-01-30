package piscine

func ToLower(s string) string {
	bytes := []byte(s)
	for i, b := range bytes {
		if b >= 'A' && b <= 'Z' {
			bytes[i] = b + 32
		}
	}
	return string(bytes)
}
