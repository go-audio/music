package theory

// more progressions available there: https://en.wikipedia.org/wiki/List_of_chord_progressions

var MajorProgressions = [][]int{
	{0, 3},
	{3, 0},
	{5, 3, 0, 4},
	{0, 0, 3, 5},
	{5, 4, 3, 2},
	{2, 4},
	{0, 5, 3, 4}, // https://en.wikipedia.org/wiki/50s_progression
	{0, 4, 5, 3},
	{2, 4, 5, 3},
	{4, 2, 5, 0},
	{2, 0, 4, 5},
	{0, 2, 3, 5},
	{0, 3, 4, 3},
	{0, 3, 2, 5},
	{0, 5, 3, 2},
	{0, 3, 0, 4},
	{5, 4, 0, 3},
	{4, 3, 4, 0},
	{3, 4, 2, 0},
	{3, 4, 2, 3},
}
var MinorProgressions = [][]int{
	{0, 3},
	{3, 0},
	{5, 3, 0, 4},
	{0, 0, 3, 5},
	{5, 4, 3, 2},
	{2, 4},
	{0, 5, 3, 4},
	{0, 4, 5, 3},
	{4, 2, 5, 0},
	{2, 0, 4, 5},
	{0, 2, 3, 5},
	{0, 3, 4, 3},
	{0, 3, 2, 5},
	{0, 5, 3, 2},
	{0, 3, 0, 4},
	{3, 4, 0, 0},
	{5, 4, 0, 2},
	{0, 5, 2, 2},
	{3, 0, 3, 4},
	{4, 3, 4, 0},
	{4, 3, 5, 4},
	{0, 4, 3, 4},
}

/* RomanNumerals are the roman number representation of the chords based on their mode/scale.

Major chord: I, II, III...
Minor chord: i, ii, iii...
Augmented chord: I+, II+, III+...
Diminished chord: vi°, vii°...
Half-diminished chord: viiØ7...
Extended chords: ii7, V9, V13...
Altered tones or chords: #iv, ii#7
*/
var RomanNumerals = map[ScaleName][]string{
	MajorScale:           {"I", "ii", "iii", "IV", "V", "vi", "vii°"},
	DorianScale:          {"i", "ii", "bIII", "IV", "v", "vi°", "bVII"},
	PhrygianScale:        {"i", "bII", "bIII", "iv", "v°", "bVI", "bvii"},
	LydianScale:          {"I", "II", "iii", "#iv°", "V", "vi", "vii"},
	MixolydianScale:      {"I", "ii", "iii°", "IV", "v", "vi", "bVII"},
	MinorPentatonicScale: {"i", "III", "iv", "v", "VII"},
	HarmonicMinorScale:   {"i", "ii°", "III+", "iv", "V", "VI", "bvii°"},
	MelodicMinorScale:    {"i", "ii", "III+", "IV", "V", "bvi°", "bvii°"},
	NaturalMinorScale:    {"i", "ii°", "III", "iv", "v", "VI", "VII"},
	MajorPentatonicScale: {"i", "ii", "iii", "V", "vi"},
	LocrianScale:         {"i°", "bII", "biii", "iv", "bV", "bVI", "bvii"},
	BluesScale:           {"i°", "bIII", "IV", "bV", "V", "bVII", "I"},
}
