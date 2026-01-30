package piscine

func Swap(a *int, b *int) {
	temp := *a // save value of a
	*a = *b    // put b’s value into a
	*b = temp  // put saved value into b
}
