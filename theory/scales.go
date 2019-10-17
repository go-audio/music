package theory

import (
	"github.com/go-audio/midi"
)

// ScaleDefinition defines a scale name, is it a modal and if yes, it's parent mode.
// Also holds the scale breakdown in half-tones and caches the intervals.
// Doesn't support other systems than western for now.
type ScaleDefinition struct {
	Name        ScaleName
	isModal     bool
	ParentScale ScaleName
	HalfSteps   []int
	_scaleNotes []int // cache
}

// NotesInScale returns the index 0 base 12 notes/keys within the scale
func (def ScaleDefinition) NotesInScale() []int {
	if len(def._scaleNotes) > 0 {
		return def._scaleNotes
	}
	def._scaleNotes = []int{}
	for i := 0; i < 12; i++ {
		if def.InScale[i] {
			def._scaleNotes = append(def._scaleNotes, i)
		}
	}
	return def._scaleNotes
}

// ScaleDefinitions is a type representing slice of scale definitions
type ScaleDefinitions []ScaleDefinition

// ScaleName is the English name of the scale
type ScaleName string

// Scales Below are grouped by modes derivated from ionian, harmonic minor, melodic minor
// and the ones that are outside of modal system.
// Using S as Sharp (#) indicator for now.
const (
	Ionian     ScaleName = "Ionian"
	Dorian     ScaleName = "Dorian"
	Phrygian   ScaleName = "Phrygian"
	Lydian     ScaleName = "Lydian"
	Mixolydian ScaleName = "Mixolydian"
	Aeolian    ScaleName = "Aeolian"
	Locrian    ScaleName = "Locrian"

	HarmonicMinor    ScaleName = "HarmonicMinor"
	LocrianS6        ScaleName = "LocrianS6"
	IonianS5         ScaleName = "IonianS5"
	DorianS4         ScaleName = "DorianS4"
	PhrygianDominant ScaleName = "PhrygianDominant"
	LydianS2         ScaleName = "LydianS2"
	SuperLocrian     ScaleName = "SuperLocrian"

	MelodicMinor    ScaleName = "MelodicMinor"
	Dorianb2        ScaleName = "Dorianb2"
	LydianAugmented ScaleName = "LydianAugmented"
	LydianDominant  ScaleName = "LydianDominant"
	Mixolydianb6    ScaleName = "Mixolydianb6"
	Aeolianb5       ScaleName = "Aeolianb5"
	Altered         ScaleName = "Altered"

	WholeTone       ScaleName = "WholeTone"
	Blues           ScaleName = "Blues"
	MinorPentatonic ScaleName = "MinorPentatonic"
	MajorPentatonic ScaleName = "MajorPentatonic"
	Diminished      ScaleName = "Diminished"
	MajorBebop      ScaleName = "MajorBebop"
	MinorBebop      ScaleName = "MinorBebop"
	DominantBebop   ScaleName = "DominantBebop"
	Arabic          ScaleName = "Arabic"
)

var (
	// ScaleDefs list all scales intervals
	ScaleDefs = []ScaleDefinition{
		0: {Name: Ionian,
			isModal:     true,
			ParentScale: nil,
			HalfSteps:   []int{2, 2, 1, 2, 2, 2, 1},
		},
		1: {Name: Dorian,
			isModal:     true,
			ParentScale: Ionian,
			HalfSteps:   []int{2, 1, 2, 2, 2, 1, 2},
		},
		2: {Name: Phrygian,
			isModal:     true,
			ParentScale: Ionian,
			HalfSteps:   []int{1, 2, 2, 2, 1, 2, 2},
		},
		3: {Name: Lydian,
			isModal:     true,
			ParentScale: Ionian,
			HalfSteps:   []int{2, 2, 2, 1, 2, 2, 1},
		},
		4: {Name: Mixolydian,
			isModal:     true,
			ParentScale: Ionian,
			HalfSteps:   []int{2, 2, 1, 2, 2, 2, 2},
		},
		5: {Name: Aeolian,
			isModal:     true,
			ParentScale: Ionian,
			HalfSteps:   []int{2, 1, 2, 2, 2, 1, 2},
		},
		6: {Name: Locrian,
			isModal:     true,
			ParentScale: Ionian,
			HalfSteps:   []int{1, 2, 2, 1, 2, 2, 2},
		},
		7: {Name: HarmonicMinor,
			isModal:     true,
			ParentScale: nil,
			HalfSteps:   []int{2, 1, 2, 2, 1, 3, 1},
		},
		8: {Name: LocrianS6,
			isModal:     true,
			ParentScale: HarmonicMinor,
			HalfSteps:   []int{1, 2, 2, 1, 3, 1, 2},
		},
		9: {Name: IonianS5,
			isModal:     true,
			ParentScale: HarmonicMinor,
			HalfSteps:   []int{2, 2, 1, 3, 1, 2, 1},
		},
		10: {Name: DorianS4,
			isModal:     true,
			ParentScale: HarmonicMinor,
			HalfSteps:   []int{2, 1, 3, 1, 2, 1, 2},
		},
		11: {Name: PhrygianDominant,
			isModal:     true,
			ParentScale: HarmonicMinor,
			HalfSteps:   []int{1, 3, 1, 2, 1, 2, 2},
		},
		12: {Name: LydianS2,
			isModal:     true,
			ParentScale: HarmonicMinor,
			HalfSteps:   []int{3, 1, 2, 1, 2, 2, 1},
		},
		13: {Name: SuperLocrian,
			isModal:     true,
			ParentScale: HarmonicMinor,
			HalfSteps:   []int{1, 2, 1, 2, 2, 1, 3},
		},
		14: {Name: MelodicMinor,
			isModal:     true,
			ParentScale: nil,
			HalfSteps:   []int{2, 1, 2, 2, 2, 2, 1},
		},
		15: {Name: Dorianb2,
			isModal:     true,
			ParentScale: MelodicMinor,
			HalfSteps:   []int{1, 2, 2, 2, 2, 1, 2},
		},
		16: {Name: LydianAugmented,
			isModal:     true,
			ParentScale: MelodicMinor,
			HalfSteps:   []int{2, 2, 2, 2, 1, 2, 1},
		},
		17: {Name: LydianDominant,
			isModal:     true,
			ParentScale: MelodicMinor,
			HalfSteps:   []int{2, 2, 2, 1, 2, 1, 2},
		},
		18: {Name: Mixolydianb6,
			isModal:     true,
			ParentScale: MelodicMinor,
			HalfSteps:   []int{2, 2, 1, 2, 1, 2, 2},
		},
		19: {Name: Aeolianb5,
			isModal:     true,
			ParentScale: MelodicMinor,
			HalfSteps:   []int{2, 1, 2, 1, 2, 2, 2},
		},
		20: {Name: Altered,
			isModal:     true,
			ParentScale: MelodicMinor,
			HalfSteps:   []int{1, 2, 1, 2, 2, 2, 2},
		},
		21: {Name: WholeTone,
			isModal:     false,
			ParentScale: nil,
			HalfSteps:   []int{2, 2, 2, 2, 2, 2},
		},
		22: {Name: Blues,
			isModal:     false,
			ParentScale: nil,
			HalfSteps:   []int{3, 2, 1, 1, 3, 2},
		},
		23: {Name: MinorPentatonic,
			isModal:     false,
			ParentScale: nil,
			HalfSteps:   []int{3, 2, 2, 3, 2},
		},
		24: {Name: MajorPentatonic,
			isModal:     false,
			ParentScale: nil,
			HalfSteps:   []int{2, 2, 3, 2, 3},
		},
		25: {Name: Diminished,
			isModal:     false,
			ParentScale: nil,
			HalfSteps:   []int{3, 3, 3, 3},
		},
		26: {Name: MajorBebop,
			isModal:     false,
			ParentScale: nil,
			HalfSteps:   []int{2, 2, 1, 2, 1, 1, 2, 1},
		},
		27: {Name: MinorBebop,
			isModal:     false,
			ParentScale: nil,
			HalfSteps:   []int{2, 1, 1, 1, 2, 2, 1, 2},
		},
		28: {Name: DominantBebop,
			isModal:     false,
			ParentScale: nil,
			HalfSteps:   []int{2, 2, 1, 2, 2, 1, 1, 1},
		},
		29: {Name: Arabic,
			isModal:     false,
			ParentScale: nil,
			HalfSteps:   []int{1, 3, 1, 2, 1, 3, 1},
		},
	}

	// ScaleDefMap is a map of the available scales
	ScaleDefMap = map[ScaleName]ScaleDefinition{
		Ionian:     ScaleDefs[0],
		Dorian:     ScaleDefs[1],
		Phrygian:   ScaleDefs[2],
		Lydian:     ScaleDefs[3],
		Mixolydian: ScaleDefs[4],
		Aeolian:    ScaleDefs[5],
		Locrian:    ScaleDefs[6],

		HarmonicMinor:    ScaleDefs[7],
		LocrianS6:        ScaleDefs[8],
		IonianS5:         ScaleDefs[9],
		DorianS4:         ScaleDefs[10],
		PhrygianDominant: ScaleDefs[11],
		LydianS2:         ScaleDefs[12],
		SuperLocrian:     ScaleDefs[13],

		MelodicMinor:    ScaleDefs[14],
		Dorianb2:        ScaleDefs[15],
		LydianAugmented: ScaleDefs[16],
		LydianDominant:  ScaleDefs[17],
		Mixolydianb6:    ScaleDefs[18],
		Aeolianb5:       ScaleDefs[19],
		Altered:         ScaleDefs[20],

		WholeTone:       ScaleDefs[21],
		Blues:           ScaleDefs[22],
		MinorPentatonic: ScaleDefs[23],
		MajorPentatonic: ScaleDefs[24],
		Diminished:      ScaleDefs[25],
		MajorBebop:      ScaleDefs[26],
		MinorBebop:      ScaleDefs[27],
		DominantBebop:   ScaleDefs[28],
		Arabic:          ScaleDefs[29],
	}
)

// ScaleNotes returns the notes in the scale. The return data contains the
// note numbers (0-11) and the English musical notes
func ScaleNotes(tonic string, scale ScaleName) ([]int, []string) {
	k := midi.KeyInt(tonic, 0) % 12
	scaleKeys := []int{k}
	for _, hs := range ScaleDefMap[scale].HalfSteps {
		k += hs
		scaleKeys = append(scaleKeys, k%12)
	}
	notes := []string{}
	for _, k := range scaleKeys {
		notes = append(notes, midi.Notes[k%12])
	}
	return scaleKeys, notes
}
