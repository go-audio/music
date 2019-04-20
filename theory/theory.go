package theory

var scaleDegreeNames = []string{
	"Tonic",
	"Supertonic",
	"Mediant",
	"Subdominant",
	"Dominant",
	"Submediant",
	"Leading tone/Subtonic", // Leading tone (in Major scale) / Subtonic (in Natural Minor Scale)
}

// ScaleDegreeName returns the name of the position of a particuliar note on a
// scale. The position is expected in index 0 and between 0 and 6.
// https://en.wikipedia.org/wiki/Degree_(music)
func ScaleDegreeName(pos int) string {
	if pos < 0 || pos > 7 {
		return "Out of scale"
	}
	if pos == 7 {
		return "Tonic (octave)"
	}
	return scaleDegreeNames[pos]
}

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
