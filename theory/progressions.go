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
Minor chord: Im, IIm, IIIm...
Augmented chord: I(#5), II(#5), III(#5)...
Diminished chord: VI(b5), VI(b5)...
Half-diminished chord: VIIm7b5...
Extended chords: IIm7, Imaj7, V9, V13...
Altered tones or chords: IIm7b9b11

For now will follow only the modal ("church") scales. No extended chords (root + third + fifth).
*/
var RomanNumerals = map[ScaleName][]string{
	Ionian:     {"I", "IIm", "IIIm", "IV", "V", "VIm", "VIIm(b5)"},
	Dorian:     {"Im", "IIm", "bIII", "IV", "V", "VIm(b5)", "bVII"},
	Phrygian:   {"Im", "bII", "bIII", "IVm", "VIm(b5)", "bVI", "bVIIm"},
	Lydian:     {"I", "II", "IIIm", "#IVm(b5)", "V", "VIm", "VIIm"},
	Mixolydian: {"I", "IIm", "IIIm(b5)", "IV", "Vm", "VIm", "bVII"},
	Aeolian:    {"Im", "IIm(b5)", "III", "IVm", "Vm", "VI", "VII"},
	Locrian:    {"Im(b5)", "bII", "bIIIm", "IVm", "bV", "bVI", "bVIIm"},
}
