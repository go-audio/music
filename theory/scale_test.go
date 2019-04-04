package theory

import (
	"reflect"
	"testing"

	"github.com/go-audio/midi"
)

func TestScale_String(t *testing.T) {
	type fields struct {
		Root int
		Def  ScaleDefinition
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "C Major",
			fields: fields{Root: 60, Def: ScaleDefMap[MajorScale]},
			want:   "C Major",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scale{
				Root: tt.fields.Root,
				Def:  tt.fields.Def,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("Scale.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEligibleScalesForNotes(t *testing.T) {
	tests := []struct {
		name    string
		notes   []int
		popular bool
		want    Scales
	}{
		{"C Major",
			[]int{
				midi.KeyInt("E", 3),
				midi.KeyInt("C", 3),
				midi.KeyInt("B", 2),
				midi.KeyInt("G", 3),
				midi.KeyInt("F", 3),
			},
			true,
			Scales{
				{Root: 0, Def: ScaleDefMap[MajorScale]},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EligibleScalesForNotes(tt.notes)
			if tt.popular {
				got = got.Popular()
			}
			for _, s := range got {
				t.Logf("%s\n", s.String())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EligibleScalesForNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScale_OffsetForNote(t *testing.T) {
	tests := []struct {
		name       string
		scale      *Scale
		inputNotes []int
		want       []int
	}{
		{
			name: "3rd octave on C Major scale",
			scale: &Scale{
				Root: midi.KeyInt("C", 3) % 12,
				Def:  ScaleDefMap[MajorScale],
			},
			inputNotes: []int{60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71},
			want:       []int{0, -1, 0, -1, 0, 0, -1, 0, -1, 0, -1, 0},
		},
		{
			name: "3rd octave on C# Major scale",
			scale: &Scale{
				Root: midi.KeyInt("C#", 3) % 12,
				Def:  ScaleDefMap[MajorScale],
			},
			inputNotes: []int{60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71},
			want:       []int{0, 0, -1, 0, -1, 0, 0, -1, 0, -1, 0, -1},
		},
		{
			name: "3rd octave on C Minor scale",
			scale: &Scale{
				Root: midi.KeyInt("C", 3) % 12,
				Def:  ScaleDefMap[NaturalMinorScale],
			},
			inputNotes: []int{60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71},
			want:       []int{0, -1, 0, 0, -1, 0, -1, 0, 0, -1, 0, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, note := range tt.inputNotes {
				if got := tt.scale.OffsetForNote(note); got != tt.want[i] {
					t.Errorf("Scale.OffsetForNote(%d) = %v, want %v", note, got, tt.want[i])
				}
			}
		})
	}
}

func TestScale_AdjustedNote(t *testing.T) {
	tests := []struct {
		name       string
		scale      *Scale
		inputNotes []int
		want       []int
	}{
		{
			name: "3rd octave on C Major",
			scale: &Scale{
				Root: midi.KeyInt("C", 3),
				Def:  ScaleDefMap[MajorScale],
			},
			inputNotes: []int{60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71},
			want:       []int{60, 60, 62, 62, 64, 65, 65, 67, 67, 69, 69, 71},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, note := range tt.inputNotes {
				if got := tt.scale.AdjustedNote(note); got != tt.want[i] {
					t.Errorf("Scale.AdjustedNote(%d) = %v, want %v", note, got, tt.want[i])
				}
			}
		})
	}
}
