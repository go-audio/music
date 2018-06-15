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
