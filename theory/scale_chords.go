package theory

// ScaleChords lists the chords formed in order
var ScaleChords = map[ScaleName][]string{
	MajorScale: []string{"maj", "min", "min", "maj", "maj", "min", "mb5"},
	// minor diminished major minor minor major major
	NaturalMinorScale: []string{"min", "mb5", "maj", "min", "min", "maj", "maj"},
}

// triad
// four note chords (7th)

var RichScaleChords = map[ScaleName][][]string{
	MajorScale: [][]string{
		{"maj", "Maj7"},
		{"min", "m7"},
		{"min", "m7"},
		{"maj", "Maj7"},
		{"maj", "7"},
		{"min", "m7"},
		{"mb5", "m7b5"},
	},
	//
	NaturalMinorScale: [][]string{
		{"min", "m7"},
		{"mb5", "m7b5"},
		{"maj", "Maj7"},
		{"min", "m7"},
		{"min", "m7"},
		{"maj", "Maj7"},
		{"maj", "7"},
	},
}
