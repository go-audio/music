package theory

import (
	"reflect"
	"testing"

	"github.com/go-audio/midi"
)

func TestChord_String(t *testing.T) {
	tests := []struct {
		name string
		root string
		keys []int
		want string
	}{
		{
			name: "unordered chord",
			keys: []int{
				midi.KeyInt("F#", 3),
				midi.KeyInt("A", 3),
				midi.KeyInt("D", 4),
			},
			want: `D Major - "D4, F#3, A3"`,
		},
		{
			name: "no notes",
			keys: []int{},
			want: `Unknown - ""`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chord{Keys: tt.keys}
			if out := c.String(); out != tt.want {
				t.Errorf("Expected %s but got %s", tt.want, out)
			}
		})
	}
}

func TestChord_Def(t *testing.T) {
	tests := []struct {
		name string
		keys []int
		want string
		// Yes I'm lazy and should create 2 separate tests...
		toString string
	}{
		{
			name: "unordered chord",
			keys: []int{
				midi.KeyInt("F#", 3),
				midi.KeyInt("A", 3),
				midi.KeyInt("D", 4),
			},
			want:     "D Major",
			toString: `D Major - "D4, F#3, A3"`,
		},
		{
			name: "Bm",
			keys: []int{
				midi.KeyInt("B", 2),
				midi.KeyInt("D", 3),
				midi.KeyInt("F#", 3),
			},
			want:     "B Minor",
			toString: `B Minor - "B2, D3, F#3"`,
		},
		{
			name: "Bm\n",
			keys: []int{
				midi.KeyInt("B", 2),
				midi.KeyInt("D", 3),
				midi.KeyInt("F#", 3),
			},
			want:     "B Minor",
			toString: `B Minor - "B2, D3, F#3"`,
		},
		{
			name: "F#m",
			keys: []int{
				midi.KeyInt("F#", 3),
				midi.KeyInt("A", 3),
				midi.KeyInt("C#", 4),
			},
			want:     "F# Minor",
			toString: `F# Minor - "F#3, A3, C#4"`,
		},
		{
			name: "F#m alt",
			keys: []int{
				midi.KeyInt("C#", 3),
				midi.KeyInt("F#", 3),
				midi.KeyInt("A", 3),
			},
			want:     "F# Minor",
			toString: `F# Minor - "F#3, A3, C#3"`,
		},
		{
			name: "Cmaj7",
			keys: []int{
				midi.KeyInt("C", 3),
				midi.KeyInt("E", 3),
				midi.KeyInt("G", 3),
				midi.KeyInt("B", 4),
			},
			want:     "C Major Seventh",
			toString: `C Major Seventh - "C3, E3, G3, B4"`,
		},
		{
			name: "C7",
			keys: []int{
				midi.KeyInt("C", 3),
				midi.KeyInt("E", 3),
				midi.KeyInt("G", 3),
				midi.KeyInt("A#", 4),
			},
			want:     "C Seventh",
			toString: `C Seventh - "C3, E3, G3, A#4"`,
		},
		{
			name: "Cmin13",
			keys: []int{
				midi.KeyInt("C", 3),
				midi.KeyInt("D#", 3),
				midi.KeyInt("G", 3),
				midi.KeyInt("A#", 4),
				midi.KeyInt("D", 4),
				midi.KeyInt("A", 5),
			},
			want:     "C Minor Thirteenth",
			toString: `C Minor Thirteenth - "C3, D#3, G3, A#4, D4, A5"`,
		},
		{
			name: "Not enough keys for a chord",
			keys: []int{
				midi.KeyInt("C#", 3),
			},
			want:     "Unknown",
			toString: `Unknown - "C#3"`,
		},
		{
			name: "Not a chord",
			keys: []int{
				midi.KeyInt("C#", 3),
				midi.KeyInt("D", 3),
			},
			want:     "Unknown",
			toString: `Unknown - "C#3, D3"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chord{
				Keys: tt.keys,
			}
			got := c.Def().String()
			if got != tt.want {
				t.Errorf("Expected chord name: %s, got %s", tt.want, got)
			}
			if stringConv := c.String(); stringConv != tt.toString {
				t.Errorf("The string conversion failed, expected %s, got %s", tt.toString, stringConv)
			}

		})
	}
}

func TestChord_Intervals(t *testing.T) {
	tests := []struct {
		name  string
		chord bool
		keys  []int
		want  []uint
	}{
		{
			name: "across octaves",
			keys: []int{
				midi.KeyInt("B", 2),
				midi.KeyInt("D", 3),
				midi.KeyInt("F#", 3),
			},
			want: []uint{3, 4},
		},
		{
			name: "another minor",
			keys: []int{
				midi.KeyInt("F#", 3),
				midi.KeyInt("A", 3),
				midi.KeyInt("C#", 4),
			},
			want: []uint{3, 4},
		},
		{
			name: "C13",
			keys: []int{
				midi.KeyInt("C", 3),
				midi.KeyInt("D#", 3),
				midi.KeyInt("G", 3),
				midi.KeyInt("A#", 4),
				midi.KeyInt("D", 4),
				midi.KeyInt("F", 4),
				midi.KeyInt("A", 4),
			},
			want: []uint{3, 4, 3, 4, 3, 4},
		},
		{
			name:  "duplicate notes in a 5th chord",
			chord: true,
			keys: []int{
				midi.KeyInt("E", 2),
				midi.KeyInt("B", 2),
				midi.KeyInt("E", 3),
			},
			want: []uint{7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chord{Keys: tt.keys}
			if got := c.Intervals(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HomogramDistItem.Intervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewChordFromAbbrev(t *testing.T) {
	tests := []struct {
		name string
		want *Chord
	}{
		{name: "Bmin",
			want: &Chord{
				Keys: []int{
					midi.KeyInt("B", 0),
					midi.KeyInt("D", 1),
					midi.KeyInt("F#", 1),
				},
			},
		},
		{name: "Dmaj",
			want: &Chord{
				Keys: []int{
					midi.KeyInt("D", 0),
					midi.KeyInt("F#", 0),
					midi.KeyInt("A", 0),
				},
			},
		},
		{name: "F#min",
			want: &Chord{
				Keys: []int{
					midi.KeyInt("F#", 0),
					midi.KeyInt("A", 0),
					midi.KeyInt("C#", 1),
				},
			},
		},
		{name: "Emaj",
			want: &Chord{
				Keys: []int{
					midi.KeyInt("E", 0),
					midi.KeyInt("G#", 0),
					midi.KeyInt("B", 0),
				},
			},
		},
		{name: "E5",
			want: &Chord{
				Keys: []int{
					midi.KeyInt("E", 0),
					midi.KeyInt("B", 0),
				},
			},
		},
		{name: "Matt", want: nil},
		{name: "CMajor", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChordFromAbbrev(tt.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChordFromAbbrev() = %v, want %v", got, tt.want)
			}
		})
	}
}
