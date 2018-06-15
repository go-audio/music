package theory

import (
	"fmt"
	"sort"

	"github.com/go-audio/midi"
)

// Scale in this context is a musical scale including a root/tonic.
type Scale struct {
	Root int
	Def  ScaleDefinition
}

func (s *Scale) String() string {
	return fmt.Sprintf("%s %s", midi.Notes[s.Root%12], s.Def.Name)
}

// Scales is a slice of scales
type Scales []Scale

// Popular filter down to only return the popular scales
func (scales Scales) Popular() Scales {
	out := Scales{}
	for _, scale := range scales {
		if scale.Def.Popular {
			out = append(out, scale)
		}
	}
	return out
}

// EligibleScalesForNotes returns a slice of scales that would satisfy the passed notes.
func EligibleScalesForNotes(notes []int) Scales {
	// remove duplicates and -1s
	uNoteMap := map[int]bool{}
	for _, n := range notes {
		if n >= 0 {
			uNoteMap[n%12] = true
		}
	}
	uNotes := make([]int, len(uNoteMap))

	var i int
	for k := range uNoteMap {
		uNotes[i] = k
		i++
	}
	scales := map[string]Scale{}
	// compare our list of unique notes to each scale to see if they are compatible
	// we need to test each scale in each unique notes we have.
	for _, root := range uNotes {
		for _, def := range ScaleDefs {
			notes, _ := ScaleNotes(midi.Notes[root%12], def.Name)
			if intSliceIncludesOther(notes, uNotes) {
				scale := Scale{Root: root % 12, Def: def}
				scales[scale.String()] = scale
			}
		}
	}

	// force a sorted list so we get a consistent output
	scaleList := make([]Scale, len(scales))
	scaleNames := make([]string, len(scales))
	i = 0
	for name, _ := range scales {
		scaleNames[i] = name
		i++
	}
	sort.Strings(scaleNames)
	for i, name := range scaleNames {
		scaleList[i] = scales[name]
	}

	sort.Slice(scaleList, func(i, j int) bool { return scaleList[i].String() < scaleList[i].String() })
	return scaleList
}
