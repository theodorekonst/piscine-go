package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	asc := true
	desc := true

	for i := 0; i < len(a)-1; i++ {
		comp := f(a[i], a[i+1])
		if comp > 0 {
			// a[i] > a[i+1], άρα δεν είναι αυστηρά αύξουσα
			asc = false
		}
		if comp < 0 {
			// a[i] < a[i+1], άρα δεν είναι αυστηρά φθίνουσα
			desc = false
		}
	}

	return asc || desc
}
