package piscine

func TrimAtoi(s string) int {
	sign := 1
	num := 0
	foundDigit := false

	for _, r := range s {
		if r == '-' && !foundDigit {
			sign = -1
		}
		if r >= '0' && r <= '9' {
			foundDigit = true
			num = num*10 + int(r-'0')
		}
	}
	return sign * num
}
