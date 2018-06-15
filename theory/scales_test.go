package theory

import (
	"reflect"
	"testing"
)

func TestScaleNotes(t *testing.T) {
	tests := []struct {
		name      string
		tonic     string
		scale     ScaleName
		wantKeys  []int
		wantNames []string
	}{
		{
			name: "C Major", tonic: "c", scale: MajorScale,
			wantKeys:  []int{0, 2, 4, 5, 7, 9, 11},
			wantNames: []string{`C`, `D`, `E`, `F`, `G`, `A`, `B`},
		},
		{
			name: "C melodic Minor", tonic: "C", scale: MelodicMinorScale,
			wantKeys:  []int{0, 2, 3, 5, 7, 9, 11},
			wantNames: []string{`C`, `D`, `D#`, `F`, `G`, `A`, `B`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys, notes := ScaleNotes(tt.tonic, tt.scale)
			if !reflect.DeepEqual(keys, tt.wantKeys) {
				t.Errorf("ScaleNotes() keys = %v, want %v", keys, tt.wantKeys)
			}
			if !reflect.DeepEqual(notes, tt.wantNames) {
				t.Errorf("ScaleNotes() notes = %v, want %v", notes, tt.wantNames)
			}
		})
	}
}
