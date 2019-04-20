package theory

import (
	"strconv"
)

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

// OrdinalPositionName returns the ordinal name of an index 0 position.
// 1st, 2nd, 3rd, 4th, 5th etc.
func OrdinalPositionName(pos int) string {
	cardinal := pos + 1
	if cardinal < 1 {
		return "invalid position"
	}
	suffix := "th"
	switch cardinal % 10 {
	case 1:
		if cardinal%100 != 11 {
			suffix = "st"
		}
	case 2:
		if cardinal%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if cardinal%100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(cardinal) + suffix
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
