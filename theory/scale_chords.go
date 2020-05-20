package theory

// ScaleChords lists the chords formed in order, the first entry is for triads,
// second 7th and last 9th
var ScaleChords = map[ScaleName][][]string{
	Ionian: [][]string{
		{"maj", "Maj7", "Maj9"},
		{"min", "m7", "m9"},
		{"min", "m7", "m9"},
		{"maj", "Maj7", "Maj9"},
		{"maj", "7", "9"},
		{"min", "m7", "m9"},
		{"mb5", "m7b5", "m9b5"},
	},
	//
	Aeolian: [][]string{
		{"min", "m7", "m9"},
		{"mb5", "m7b5", "m9b5"},
		{"maj", "Maj7", "Maj9"},
		{"min", "m7", "m9"},
		{"min", "m7", "m9"},
		{"maj", "Maj7", "Maj9"},
		{"maj", "7", "9"},
	},
}
