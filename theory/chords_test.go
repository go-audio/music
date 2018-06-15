package theory

import (
	"reflect"
	"testing"

	"github.com/go-audio/midi"
)

func TestChords_String(t *testing.T) {
	tests := []struct {
		name   string
		chords Chords
		want   string
	}{
		{
			name: "basic",
			chords: Chords{
				{
					Keys: []int{
						midi.KeyInt("B", 2),
						midi.KeyInt("D", 3),
						midi.KeyInt("F#", 3),
					},
				},
				{
					Keys: []int{
						midi.KeyInt("D", 4),
						midi.KeyInt("F#", 3),
						midi.KeyInt("A", 3),
					},
				},
				{
					Keys: []int{
						midi.KeyInt("F#", 3),
						midi.KeyInt("A", 3),
						midi.KeyInt("C#", 4),
					},
				},
				{
					Keys: []int{
						midi.KeyInt("E", 3),
						midi.KeyInt("G#", 3),
						midi.KeyInt("B", 3),
					},
				},
			},
			want: "Bmin,Dmaj,F#min,Emaj",
		},
		{
			name: "1 chord",
			chords: Chords{
				{
					Keys: []int{
						midi.KeyInt("B", 2),
						midi.KeyInt("D", 3),
						midi.KeyInt("F#", 3),
					},
				},
			},
			want: "Bmin",
		},
		{
			name: "E5 2 notes",
			chords: Chords{
				{
					Keys: []int{
						midi.KeyInt("E", 0),
						midi.KeyInt("B", 0),
					},
				},
			},
			want: "E5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.chords.String(); got != tt.want {
				t.Errorf("Chords.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChords_Uniques(t *testing.T) {
	tests := []struct {
		name   string
		chords Chords
		want   string
	}{
		{
			name: "duplicates",
			chords: Chords{
				{
					Keys: []int{
						midi.KeyInt("B", 2),
						midi.KeyInt("D", 3),
						midi.KeyInt("F#", 3),
					},
				},
				{
					Keys: []int{
						midi.KeyInt("D", 4),
						midi.KeyInt("F#", 3),
						midi.KeyInt("A", 3),
					},
				},
				{
					Keys: []int{
						midi.KeyInt("D", 4),
						midi.KeyInt("F#", 4),
						midi.KeyInt("A", 4),
						midi.KeyInt("F#", 3),
						midi.KeyInt("A", 3),
					},
				},
			},
			want: "Dmaj,Bmin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// we have to sort otherwise the tests will randomly fail due to ordering
			if got := tt.chords.Uniques().SortedOnRoots().String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chords.Uniques() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChords_SortedOnRoots(t *testing.T) {
	tests := []struct {
		name   string
		chords Chords
		want   string
	}{
		{
			name: "duplicates",
			chords: Chords{
				{
					Keys: []int{
						midi.KeyInt("E", 3),
						midi.KeyInt("G#", 3),
						midi.KeyInt("B", 3),
					},
				},
				{
					Keys: []int{
						midi.KeyInt("B", 2),
						midi.KeyInt("D", 3),
						midi.KeyInt("F#", 3),
					},
				},
				{
					Keys: []int{
						midi.KeyInt("D", 4),
						midi.KeyInt("F#", 3),
						midi.KeyInt("A", 3),
					},
				},
				{
					Keys: []int{
						midi.KeyInt("D", 4),
						midi.KeyInt("F#", 4),
						midi.KeyInt("A", 4),
						midi.KeyInt("F#", 3),
						midi.KeyInt("A", 3),
					},
				},
			},
			want: "Dmaj,Emaj,Bmin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.chords.Uniques().SortedOnRoots().String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chords.Uniques() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestChords_EligibleScales(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		chords Chords
// 		root   string
// 		want   string
// 	}{
// 		{
// 			name: "C Major",
// 			root: "c",
// 			chords: Chords{
// 				{Keys: []int{
// 					midi.KeyInt("C", 3),
// 					midi.KeyInt("E", 3),
// 					midi.KeyInt("G", 3),
// 				}},
// 				{Keys: []int{
// 					midi.KeyInt("E", 3),
// 					midi.KeyInt("G#", 3),
// 					midi.KeyInt("B", 3),
// 				}},
// 				{Keys: []int{
// 					midi.KeyInt("G", 3),
// 					midi.KeyInt("B", 3),
// 					midi.KeyInt("D", 3),
// 				}},
// 			},
// 			want: "Major",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			scales := tt.chords.EligibleScales()
// 			scaleNames := []string{}
// 			for _, s := range scales {
// 				scaleNames = append(scaleNames, string(s.Name))
// 			}
// 			if got := strings.Join(scaleNames, ","); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Chords.EligibleScales() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestChords_ToFromBytes(t *testing.T) {
	// testing byte conversion round trip
	tests := []struct {
		name   string
		chords Chords
	}{
		{
			name: "Get Lucky",
			chords: Chords{
				&Chord{
					Keys: []int{
						midi.KeyInt("B", 0),
						midi.KeyInt("D", 1),
						midi.KeyInt("F#", 1),
					},
				},
				&Chord{
					Keys: []int{
						midi.KeyInt("D", 0),
						midi.KeyInt("F#", 0),
						midi.KeyInt("A", 0),
					},
				},
				&Chord{
					Keys: []int{
						midi.KeyInt("F#", 0),
						midi.KeyInt("A", 0),
						midi.KeyInt("C#", 1),
					},
				},
				&Chord{
					Keys: []int{
						midi.KeyInt("E", 0),
						midi.KeyInt("G#", 0),
						midi.KeyInt("B", 0),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, gotDict := tt.chords.ToBytes()
			gotChords := ChordsFromBytes(gotData, gotDict)
			if gotChords.String() != tt.chords.String() {
				t.Errorf("ChordsFromBytes(chords.ToBytes()) got = %v, want %v", gotChords, tt.chords)
			}
			if len(gotChords) == len(tt.chords) {
				for i, c := range gotChords {
					if c.String() != tt.chords[i].String() {
						t.Errorf("Chord [%d] got: %v wanted %v\n", i, c, tt.chords[i])
					}
				}
			}
		})
	}
}

func TestChords_ProgressionDesc(t *testing.T) {
	tests := []struct {
		name   string
		chords Chords
		want   string
	}{
		{
			name: "get lucky",
			chords: Chords{
				&Chord{
					Keys: []int{
						midi.KeyInt("B", 0),
						midi.KeyInt("D", 1),
						midi.KeyInt("F#", 1),
					},
				},
				&Chord{
					Keys: []int{
						midi.KeyInt("D", 0),
						midi.KeyInt("F#", 0),
						midi.KeyInt("A", 0),
					},
				},
				&Chord{
					Keys: []int{
						midi.KeyInt("F#", 0),
						midi.KeyInt("A", 0),
						midi.KeyInt("C#", 1),
					},
				},
				&Chord{
					Keys: []int{
						midi.KeyInt("E", 0),
						midi.KeyInt("G#", 0),
						midi.KeyInt("B", 0),
					},
				},
			},
			want: "B D F# E \n" +
				`i° bIII V IV  in B Blues
i III v iv  in B Minor Pentatonic
i III v iv  in B Natural Minor
vi I iii ii  in D Major
v VII ii° i  in E Natural Minor
iv VI i VII  in F# Natural Minor
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.chords.ProgressionDesc()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chords.ProgressionDesc() =\n%#v, want\n%#v", got, tt.want)
			}
		})
	}
}
