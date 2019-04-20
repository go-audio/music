package theory

import (
	"fmt"
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

func TestOrdinalPositionName(t *testing.T) {
	tests := []struct {
		pos  int
		want string
	}{
		{0, "1st"},
		{1, "2nd"},
		{2, "3rd"},
		{3, "4th"},
		{9, "10th"},
		{10, "11th"},
		{11, "12th"},
		{12, "13th"},
		{99, "100th"},
		{101, "102nd"},
		{102, "103rd"},
		{-1, "invalid position"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.pos), func(t *testing.T) {
			if got := OrdinalPositionName(tt.pos); got != tt.want {
				t.Errorf("OrdinalPositionName(%d) = %v, want %v", tt.pos, got, tt.want)
			}
		})
	}
}
