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

func TestScale_TriadChordForRoot(t *testing.T) {
	tests := []struct {
		name          string
		scale         *Scale
		inputNote     int
		want          *Chord
		wantChordName string
	}{
		{
			name:          "C3 in C Major",
			scale:         &Scale{Root: midi.KeyInt("C", 3) % 12, Def: ScaleDefMap[MajorScale]},
			inputNote:     midi.KeyInt("C", 3),
			want:          &Chord{Keys: []int{60, 64, 67}},
			wantChordName: "C Major",
		},
		{
			name:          "C3 in C Minor",
			scale:         &Scale{Root: midi.KeyInt("C", 3) % 12, Def: ScaleDefMap[NaturalMinorScale]},
			inputNote:     midi.KeyInt("C", 3),
			want:          &Chord{Keys: []int{60, 63, 67}},
			wantChordName: "C Minor",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.scale.TriadChordForRoot(tt.inputNote)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale.TriadChordForRoot() = %v, want %v", got, tt.want)
			}
			if got.Def().String() != tt.wantChordName {
				t.Errorf("Scale.TriadChordForRoot() = %v, want %v", got.Def().String(), tt.wantChordName)
			}
		})
	}
}

func TestScale_SeventhChordForRoot(t *testing.T) {
	tests := []struct {
		name          string
		scale         *Scale
		inputNote     int
		want          *Chord
		wantChordName string
	}{
		{
			name:          "C3 in C Major",
			scale:         &Scale{Root: midi.KeyInt("C", 3) % 12, Def: ScaleDefMap[MajorScale]},
			inputNote:     midi.KeyInt("C", 3),
			want:          &Chord{Keys: []int{60, 64, 67, 71}},
			wantChordName: "C Major Seventh",
		},
		{
			name:          "C3 in C Minor",
			scale:         &Scale{Root: midi.KeyInt("C", 3) % 12, Def: ScaleDefMap[NaturalMinorScale]},
			inputNote:     midi.KeyInt("C", 3),
			want:          &Chord{Keys: []int{60, 63, 67, 70}},
			wantChordName: "C Minor Seventh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.scale.SeventhChordForRoot(tt.inputNote)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale.SeventhChordForRoot() = %v, want %v", got, tt.want)
			}
			if got.Def().String() != tt.wantChordName {
				t.Errorf("Scale.SeventhChordForRoot() = %v, want %v", got.Def().String(), tt.wantChordName)
			}
		})
	}
}

func TestScale_NinthChordForRoot(t *testing.T) {
	tests := []struct {
		name          string
		scale         *Scale
		inputNote     int
		want          *Chord
		wantChordName string
	}{
		{
			name:          "C3 in C Major",
			scale:         &Scale{Root: midi.KeyInt("C", 3) % 12, Def: ScaleDefMap[MajorScale]},
			inputNote:     midi.KeyInt("C", 3),
			want:          &Chord{Keys: []int{60, 64, 67, 71, 74}},
			wantChordName: "C Major Ninth",
		},
		{
			name:          "C3 in C Minor",
			scale:         &Scale{Root: midi.KeyInt("C", 3) % 12, Def: ScaleDefMap[NaturalMinorScale]},
			inputNote:     midi.KeyInt("C", 3),
			want:          &Chord{Keys: []int{60, 63, 67, 70, 74}},
			wantChordName: "C Minor Ninth",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.scale.NinthChordForRoot(tt.inputNote)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale.NinthChordForRoot() = %v, want %v", got, tt.want)
			}
			if got.Def().String() != tt.wantChordName {
				t.Errorf("Scale.NinthChordForRoot() = %v, want %v", got.Def().String(), tt.wantChordName)
			}
		})
	}
}

func TestScale_Notes(t *testing.T) {
	tests := []struct {
		name string
		root int
		def  ScaleDefinition
		want []int
	}{
		{
			name: "C Major",
			root: 0,
			def:  ScaleDefMap[MajorScale],
			want: []int{0, 2, 4, 5, 7, 9, 11},
		},
		{
			name: "D# Minor",
			root: 3,
			def:  ScaleDefMap[NaturalMinorScale],
			want: []int{3, 5, 6, 8, 10, 11, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scale{
				Root: tt.root,
				Def:  tt.def,
			}
			if got := s.Notes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale.Notes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScale_IndexOfNote(t *testing.T) {
	tests := []struct {
		name string
		root int
		def  ScaleDefinition
		note int
		want int
	}{
		{
			name: "D# 3 in D# Minor",
			root: midi.KeyInt("D#", 0),
			def:  ScaleDefMap[NaturalMinorScale],
			note: midi.KeyInt("D#", 3),
			want: 0, // index 0
		},
		{
			name: "out of scale key in D# Minor",
			root: midi.KeyInt("D#", 0),
			def:  ScaleDefMap[NaturalMinorScale],
			note: midi.KeyInt("C", 3),
			want: -1, // not there
		},
		{
			name: "F# in D# Minor",
			root: midi.KeyInt("D#", 0),
			def:  ScaleDefMap[NaturalMinorScale],
			note: midi.KeyInt("F#", 3),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scale{
				Root: tt.root,
				Def:  tt.def,
			}
			if got := s.IndexOfNote(tt.note); got != tt.want {
				t.Errorf("Scale.IndexOfNote(%d/%s) = %v, want %v", tt.note, midi.NoteToName(tt.note), got, tt.want)
			}
		})
	}
}
