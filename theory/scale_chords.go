package theory

// ScaleChords lists the chords formed in order
var ScaleChords = map[ScaleName][]string{
	MajorScale:        []string{"maj", "min", "min", "maj", "maj", "min", "mb5"},
	MelodicMinorScale: []string{"min", "min", "aug", "maj", "maj", "mb5", "mb5"},
}
