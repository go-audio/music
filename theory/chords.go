package theory

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-audio/midi"
)

// Chords is a slice of chords
type Chords []*Chord

// ToBytes compresses a slice of chords where each chord is represented by a
// byte. A dictionary is also returned so a byte can be converted back to a
// chord (but the octave and repeated notes within the chord will be lost)
func (chords Chords) ToBytes() (data []byte, dict map[byte]string) {
	// build the dictionaries
	uNames := strings.Split(chords.Uniques().String(), ",")
	reverseDict := make(map[string]byte, len(uNames))
	dict = make(map[byte]string, len(uNames))
	currentByte := byte(1)
	for _, name := range uNames {
		dict[currentByte] = name
		reverseDict[name] = currentByte
		currentByte++
	}
	// build the data
	for _, name := range strings.Split(chords.String(), ",") {
		data = append(data, reverseDict[name])
	}

	return data, dict
}

// ChordsFromBytes convert an encoded slice of bytes back to a slice of chords
// using the passed dictionary. Note that this is a lossy conversion, the octave
// of the original chords will be lost.
func ChordsFromBytes(data []byte, dict map[byte]string) Chords {
	chords := make(Chords, len(data))
	for i, b := range data {
		chords[i] = NewChordFromAbbrev(dict[b])
	}
	return chords
}

// String converts the sequence of chords in a nice, coma delimited string
// such as Bmin,Dmaj,F#min,Emaj,Bmin,Dmaj,F#min,Emaj,Bmin,Dmaj
func (chords Chords) String() string {
	if chords == nil || len(chords) < 1 {
		return ""
	}
	if len(chords) < 2 {
		if chords[0] == nil || len(chords[0].Keys) < 2 {
			return ""
		}
		return chords[0].AbbrevName()
	}
	b := strings.Builder{}
	for i := 0; i < len(chords)-1; i++ {
		if len(chords[i].Keys) > 1 {
			b.WriteString(chords[i].AbbrevName() + ",")
		}
	}
	lastIdx := len(chords) - 1
	if len(chords[lastIdx].Keys) > 1 {
		b.WriteString(chords[lastIdx].AbbrevName())
	}

	return b.String()
}

// RootNotes returns the root notes of each chord (index 0)
func (chords Chords) RootNotes() []int {
	notes := make([]int, len(chords))
	for i, c := range chords {
		notes[i] = c.Def().RootInt()
	}
	return notes
}

// EligibleScales returns a list of potentially matching scales.
// This can be used to calculate the chord progression within a scale.
func (chords Chords) EligibleScales() Scales {
	notes := chords.Uniques().SortedOnRoots().RootNotes()
	return EligibleScalesForNotes(notes)
}

// Uniques returns a copy of the unique chords contains in targetted chords.
func (chords Chords) Uniques() Chords {
	uChords := Chords{}
	mChords := map[string]*Chord{}
	var abbrevName string
	for _, c := range chords {
		abbrevName = c.AbbrevName()
		if _, ok := mChords[abbrevName]; !ok {
			mChords[abbrevName] = c
		}
	}
	for _, c := range mChords {
		uChords = append(uChords, c)
	}

	return uChords
}

// SortedOnRoots does an in place sorting of the chords based on their root
// notes. The chords are returned so the calls can be chained
func (chords Chords) SortedOnRoots() Chords {
	sort.Slice(chords, func(a, b int) bool {
		return chords[a].Def().RootInt() < chords[b].Def().RootInt()
	})
	return chords
}

// ProgressionDesc returns the roman numerals describing the figured chords.
func (chords Chords) ProgressionDesc() string {
	var out string

	notes := []int{}
	for _, c := range chords {
		if len(c.Keys) < 1 {
			continue
		}
		notes = append(notes, midi.NotesToInt[c.Def().Root])
		out += fmt.Sprintf("%s ", c.Def().Root)
	}
	out += "\n"
	//get the possible scales for the root notes
	scales := EligibleScalesForNotes(notes).Popular()
	for _, scale := range scales {
		scaleNotes, _ := ScaleNotes(midi.Notes[scale.Root%12], scale.Def.Name)
		romanScale := RomanNumerals[scale.Def.Name]
		for _, note := range notes {
			idx := sliceIndex(14, func(i int) bool { return scaleNotes[i] == note })
			if len(romanScale) > 0 {
				out += fmt.Sprintf("%s ", romanScale[idx])
			} else {
				out += fmt.Sprintf("%d ", idx)
			}
		}
		out = fmt.Sprintf("%s in %s\n", out, scale.String())
	}

	return out
}
