package theory

// ScaleChords lists the chords formed in order
var ScaleChords = map[ScaleName][]string{
	MajorScale:        []string{"maj", "min", "min", "maj", "maj", "min", "mb5"},
	MelodicMinorScale: []string{"min", "min", "aug", "maj", "maj", "mb5", "mb5"},
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
	// TODO:
	MelodicMinorScale: [][]string{
		{"min", "m7"},
		{"min", "m7b5"},
		{"aug", "Maj7"},
		{"maj", "min7"},
		{"maj", "min7"},
		{"mb5", "Maj7"},
		{"mb5", "Maj7"},
	},
}
