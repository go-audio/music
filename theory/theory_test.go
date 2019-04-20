package theory

import (
	"strconv"
	"testing"

	"github.com/go-audio/midi"
)

func keyNames(keys []int) []string {
	names := make([]string, len(keys))
	for i, k := range keys {
		names[i] = midi.NoteToName(k)
	}
	return names
}

func TestScaleDegreeName(t *testing.T) {
	tests := []struct {
		pos  int
		want string
	}{
		{-1, "Out of scale"},
		{0, "Tonic"},
		{1, "Supertonic"},
		{2, "Mediant"},
		{3, "Subdominant"},
		{4, "Dominant"},
		{5, "Submediant"},
		{6, "Leading tone/Subtonic"},
		{7, "Tonic (octave)"},
		{8, "Out of scale"},
		{99, "Out of scale"},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.pos), func(t *testing.T) {
			if got := ScaleDegreeName(tt.pos); got != tt.want {
				t.Errorf("ScaleDegreeName(%d) = %v, want %v", tt.pos, got, tt.want)
			}
		})
	}
}
