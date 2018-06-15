package theory

func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// check that a contains all elements of b (it can contain more tho)
func intSliceIncludesOther(a, b []int) bool {
	if len(b) > len(a) {
		return false
	}
	for _, n := range b {
		var isMatch bool
		for _, m := range a {
			if n == m {
				isMatch = true
				break
			}
		}
		if !isMatch {
			return false
		}
	}
	return true
}
