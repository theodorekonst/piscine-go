package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	for i := range s {
		if i+len(toFind) <= len(s) && s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
