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
		{
			name: "B Major", tonic: "b", scale: MajorScale,
			wantKeys:  []int{11, 1, 3, 4, 6, 8, 10},
			wantNames: []string{`B`, `C#`, `D#`, `E`, `F#`, `G#`, `A#`},
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

func TestScaleDefMap(t *testing.T) {
	if ScaleDefMap[MajorScale].InScale != [12]bool{true, false, true, false, true, true, false, true, false, true, false, true} {
		t.Fatalf("Expected the major scale to list notes that are in or out of the scale")
	}
}

func TestScaleDefinition_NotesInScale(t *testing.T) {
	tests := []struct {
		name string
		def  ScaleDefinition
		want []int
	}{
		{
			name: "Major",
			def:  ScaleDefMap[MajorScale],
			want: []int{0, 2, 4, 5, 7, 9, 11},
		},
		{
			name: "Minor",
			def:  ScaleDefMap[NaturalMinorScale],
			want: []int{0, 2, 3, 5, 7, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.def.NotesInScale(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScaleDefinition.NotesInScale() = %v, want %v", got, tt.want)
			}
		})
	}
}
