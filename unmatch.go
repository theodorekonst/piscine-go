package piscine

func Unmatch(a []int) int {
	for _, r := range a {
		count := 0
		for _, x := range a {
			if x == r {
				count++
			}
		}
		if count%2 == 1 {
			return r
		}
	}
	return -1
}
