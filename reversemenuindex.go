package piscine

func ReverseMenuIndex(menu []string) []string {
	newMenu := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		reverseIndex := len(menu) - 1 - i
		newMenu[i] = menu[reverseIndex]
	}
	return newMenu
}
