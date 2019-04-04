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

// OffsetForNote returns the offset to apply to an incoming note so it stays in
// scale. This is used to keep input notes within the scale (by moving the note
// to the closed lower note in scale).
func (s *Scale) OffsetForNote(note int) int {
	if s == nil {
		return 0
	}
	index := (((note % 12) + 12) - (s.Root % 12)) % 12
	if s.Def.InScale[index] {
		return 0
	}
	return -1
}

// AdjustedNote "corrects" the input note to be in scale.
func (s *Scale) AdjustedNote(note int) int {
	return note + s.OffsetForNote(note)
}

// TriadChordForRoot returns the triad chord (3 note) matching the passed root.
func (s *Scale) TriadChordForRoot(note int) *Chord {
	return s.chordForRoot(note, 3)
}

// SeventhChordForRoot returns the 7th chord (4 notes) of the scale using the
// passed root.
func (s *Scale) SeventhChordForRoot(note int) *Chord {
	return s.chordForRoot(note, 4)
}

// NinthChordForRoot returns the 7th chord (4 notes) of the scale using the
// passed root.
func (s *Scale) NinthChordForRoot(note int) *Chord {
	return s.chordForRoot(note, 5)
}

func (s *Scale) chordForRoot(note int, nbrNotesInChord uint) *Chord {
	chord := &Chord{Keys: []int{note}}
	if s == nil {
		return chord
	}
	chordsInScale := ScaleChords[s.Def.Name]
	if chordsInScale == nil {
		// unsupported scale
		return chord
	}
	// find the position of the root in the scale
	modKey := note % 12
	inScaleNotes := s.Def.NotesInScale()
	var idx int
	for i := 0; i < len(inScaleNotes); i++ {
		if modKey == inScaleNotes[i] {
			idx = i
			break
		}
	}
	var chordTypeIdx int
	if nbrNotesInChord < 4 {
		chordTypeIdx = 0 // triad
	} else if nbrNotesInChord == 4 {
		chordTypeIdx = 1 // 7th
	} else {
		chordTypeIdx = 2 // 9th
	}
	chordType := chordsInScale[idx][chordTypeIdx]
	for _, chordDef := range ChordDefs {
		if chordDef.Abbrev == chordType {
			lastNoteInChord := note
			for _, halfStep := range chordDef.HalfSteps {
				lastNoteInChord += int(halfStep)
				chord.Keys = append(chord.Keys, lastNoteInChord)
			}
			break
		}
	}
	if nbrNotesInChord < 3 {
		chord.Keys = chord.Keys[:nbrNotesInChord]
	}
	return chord
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
