package theory

import (
	"github.com/go-audio/midi"
)

// ScaleDefinition defines a scale by giving it a name and the spacing between adjacent notes on the chromatic scale.
type ScaleDefinition struct {
	// Popular indicates that the scale is commonly used
	Popular bool
	// Greek mode scale
	Greek     bool
	Name      ScaleName
	HalfSteps []int
	// InScale indicate what notes are in and which aren't
	InScale     [12]bool
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

// Popular filter down to only return the popular scales found
func (def ScaleDefinitions) Popular() ScaleDefinitions {
	out := ScaleDefinitions{}
	for _, scale := range def {
		if scale.Popular {
			out = append(out, scale)
		}
	}
	return out
}

// ScaleName is the English name of the scale
type ScaleName string

const (
	MajorScale        ScaleName = "Major"
	NaturalMinorScale ScaleName = "Natural Minor"

	HarmonicMinorScale   ScaleName = "Harmonic Minor"
	MelodicMinorScale    ScaleName = "Melodic Minor"
	WholeToneScale       ScaleName = "Whole Tone"
	DiminishedScale      ScaleName = "Diminished"
	MajorPentatonicScale ScaleName = "Major Pentatonic"
	MinorPentatonicScale ScaleName = "Minor Pentatonic"
	JapInSenScale        ScaleName = "Jap In Sen"
	MajorBebopScale      ScaleName = "Major Bebop"
	DominantBebopScale   ScaleName = "Dominant Bebop"
	BluesScale           ScaleName = "Blues"
	ArabicScale          ScaleName = "Arabic"
	EnigmaticScale       ScaleName = "Enigmatic"
	NeapolitanScale      ScaleName = "Neapolitan"
	NeapolitanMinorScale ScaleName = "Neapolitan Minor"
	HungarianMinorScale  ScaleName = "Hungarian Minor"
	DorianScale          ScaleName = "Dorian"
	PhrygianScale        ScaleName = "Phrygian"
	LydianScale          ScaleName = "Lydian"
	MixolydianScale      ScaleName = "Mixolydian"
	// LocrianScale represents the Locrian, or Hypodorian track https://en.wikipedia.org/wiki/Locrian_mode
	LocrianScale ScaleName = "Locrian"
)

var (
	// ScaleDefs list all known scales
	ScaleDefs = []ScaleDefinition{
		0: {Name: MajorScale,
			HalfSteps: []int{2, 2, 1, 2, 2, 2},
			InScale:   [12]bool{true, false, true, false, true, true, false, true, false, true, false, true},
			Popular:   true,
		},
		1: {Name: NaturalMinorScale, // AKA aeolian
			HalfSteps: []int{2, 1, 2, 2, 1, 2},
			InScale:   [12]bool{true, false, true, true, false, true, false, true, true, false, true, false},
			Popular:   true,
		},
		2:  {Name: HarmonicMinorScale, HalfSteps: []int{2, 1, 2, 2, 1, 3}},
		3:  {Name: MelodicMinorScale, HalfSteps: []int{2, 1, 2, 2, 2, 2}},
		4:  {Name: WholeToneScale, HalfSteps: []int{2, 2, 2, 2, 2}},
		5:  {Name: DiminishedScale, HalfSteps: []int{2, 1, 2, 1, 2, 1, 2}},
		6:  {Name: MajorPentatonicScale, HalfSteps: []int{2, 2, 3, 2}},
		7:  {Name: MinorPentatonicScale, HalfSteps: []int{3, 2, 2, 3}, Popular: true},
		8:  {Name: DorianScale, HalfSteps: []int{2, 1, 2, 2, 2, 1}, Greek: true},
		9:  {Name: JapInSenScale, HalfSteps: []int{1, 4, 2, 3}},
		10: {Name: MajorBebopScale, HalfSteps: []int{2, 2, 1, 2, 1, 1, 2}},
		11: {Name: DominantBebopScale, HalfSteps: []int{2, 2, 1, 2, 2, 1, 1}},
		12: {Name: BluesScale, HalfSteps: []int{3, 2, 1, 1, 3}, Popular: true},
		13: {Name: ArabicScale, HalfSteps: []int{1, 3, 1, 2, 1, 3}},
		14: {Name: EnigmaticScale, HalfSteps: []int{1, 3, 2, 2, 2, 1}},
		15: {Name: NeapolitanScale, HalfSteps: []int{1, 2, 2, 2, 2, 2}},
		16: {Name: NeapolitanMinorScale, HalfSteps: []int{1, 2, 2, 2, 1, 3}},
		17: {Name: HungarianMinorScale, HalfSteps: []int{2, 1, 3, 1, 1, 3}},
		18: {Name: PhrygianScale, HalfSteps: []int{1, 2, 2, 2, 1, 2}, Greek: true},
		19: {Name: LydianScale, HalfSteps: []int{2, 2, 2, 1, 2, 2}},
		20: {Name: MixolydianScale, HalfSteps: []int{2, 2, 1, 2, 2, 1}},
		21: {Name: LocrianScale, HalfSteps: []int{1, 2, 2, 1, 2, 2}, Greek: true},
	}

	// ScaleDefMap is a map of the available scales
	ScaleDefMap = map[ScaleName]ScaleDefinition{
		MajorScale:        ScaleDefs[0],
		NaturalMinorScale: ScaleDefs[1],
		//
		HarmonicMinorScale:   ScaleDefs[2],
		MelodicMinorScale:    ScaleDefs[3],
		WholeToneScale:       ScaleDefs[4],
		DiminishedScale:      ScaleDefs[5],
		MajorPentatonicScale: ScaleDefs[6],
		MinorPentatonicScale: ScaleDefs[7],
		DorianScale:          ScaleDefs[8],
		//
		JapInSenScale:        ScaleDefs[9],
		MajorBebopScale:      ScaleDefs[10],
		DominantBebopScale:   ScaleDefs[11],
		BluesScale:           ScaleDefs[12],
		ArabicScale:          ScaleDefs[13],
		EnigmaticScale:       ScaleDefs[14],
		NeapolitanScale:      ScaleDefs[15],
		NeapolitanMinorScale: ScaleDefs[16],
		HungarianMinorScale:  ScaleDefs[17],
		PhrygianScale:        ScaleDefs[18],
		LydianScale:          ScaleDefs[19],
		MixolydianScale:      ScaleDefs[20],
		LocrianScale:         ScaleDefs[21],
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
