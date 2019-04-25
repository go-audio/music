package theory

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/go-audio/midi"
)

// Chord represents multiple keys pressed at the same time
type Chord struct {
	// Keys are the MIDI note values for the voicing used in the chord.
	Keys []int
	// KeyIntervals are the half steps between each key, in most cases, you want to use Intervals().
	KeyIntervals     []uint
	intervalKeyCache []int
	_isSorted        bool
}

// NewChordFromAbbrev takes a chord name such as Bmin and converts it to a *Chord
// with the root key on the 0 octave.
func NewChordFromAbbrev(name string) *Chord {
	name = strings.TrimSpace(name)
	// shortest name would be a c5 or something like that
	if len(name) < 2 {
		return nil
	}
	root := []byte{name[0]}
	name = name[1:]
	if name[0] == '#' {
		root = append(root, name[0])
		name = name[1:]
	}
	var chordRef *ChordDefinition
	for _, chordDef := range ChordDefs {
		if name == chordDef.Abbrev {
			chordRef = chordDef
			break
		}
	}
	if chordRef == nil {
		return nil
	}
	chordRef.Root = string(root)
	rootInt := midi.KeyInt(chordRef.Root, 0)
	chord := &Chord{Keys: []int{rootInt}}
	for i, interv := range chordRef.HalfSteps {
		chord.Keys = append(chord.Keys, chord.Keys[i]+int(interv))
	}

	return chord
}

// AbbrevName is the abbreviated name of the chord.
func (c *Chord) AbbrevName() string {
	def := c.Def()
	if def == nil {
		return "Unknown"
	}
	if len(def.Root) > 0 {
		return fmt.Sprintf("%s%s", def.Root, def.Abbrev)
	}
	return def.Abbrev
}

func (c *Chord) String() string {
	// sort first
	def := c.Def()
	strs := make([]string, len(c.Keys))
	for i, k := range c.Keys {
		strs[i] = midi.NoteToName(k)
	}
	return fmt.Sprintf("%s - %q",
		def,
		strings.Join(strs, ", "))
}

// Def returns the matching chord definition with the root set if found. A chord
// definition with a name set to Unknown will be returned if no matches found.
// If a chord definition is found but with the notes in a different order, the
// keys will be re-ordered.
// Note that a chord could have multiple definitions, this is the best guest found.
func (c *Chord) Def() *ChordDefinition {
	if c == nil {
		return nil
	}
	var sorted bool
	// TODO: consider caching this result

	analyzedChord := c

	retries := len(analyzedChord.Keys)
	for retries > 0 {
		for _, chordDef := range ChordDefs {
			if analyzedChord.Matches(chordDef) {
				if analyzedChord != c {
					copy(c.Keys, analyzedChord.Keys)
				}
				return chordDef.WithRoot(midi.Notes[analyzedChord.Keys[0]%12])
			}
		}
		// we didn't find the chord, let's try to change the interval orders

		if len(analyzedChord.Keys) < 2 {
			break
		}
		// sort the keys and retry
		if !sorted {
			analyzedChord = analyzedChord.SortedByKeys()
			sorted = true
			continue
		}
		// we still don't have a match so we rotate the keys
		analyzedChord.Keys = append(analyzedChord.Keys[1:], analyzedChord.Keys[0])
		retries--
	}

	return &ChordDefinition{Name: "Unknown"}
}

// TODO: PossibleDefs

// SortedByKeys returns a copy of the chord but with the chord keys by pitch (lowest first)
func (c *Chord) SortedByKeys() *Chord {
	newChord := &Chord{_isSorted: true}
	// sorting the keys which might lead to issues with inversions
	sortedKeys := c.Keys
	sort.Slice(sortedKeys, func(i, j int) bool { return sortedKeys[i]%12 < sortedKeys[j]%12 })
	newChord.Keys = sortedKeys
	return newChord
	/*
		sort.Ints(sortedKeys)

		// remove duplicate notes (including those played on different octaves)
		seenKeys := map[int]bool{}
		var pitch int
		for _, k := range c.Keys {
			pitch = k % 12
			if _, ok := seenKeys[pitch]; !ok {
				seenKeys[k] = true
			}
		}

		indexOf := func(x int, list []int) int {
			for i, k := range list {
				if k%12 == x {
					return i
				}
			}
			return -1
		}

		// reorder
		newChord.Keys = make([]int, len(c.Keys))
		for pitch := range seenKeys {
			newChord.Keys[indexOf(pitch, sortedKeys)] = pitch
		}
		return newChord
	*/
}

// IsSorted checks if the keys of the chord are sorted
func (c *Chord) isSorted() bool {
	if c._isSorted {
		return true
	}
	return sort.IntsAreSorted(c.Keys)
}

// Matches compares the current chord with the passed chord.
func (c *Chord) Matches(chordDef *ChordDefinition) bool {
	if reflect.DeepEqual(chordDef.HalfSteps, c.Intervals()) {
		// confirm the root key
		for i := 1; i < len(chordDef.HalfSteps); i++ {
			if uint(c.Keys[i-1]) != chordDef.HalfSteps[i] {
				// not the root key
				continue
			}
		}
		return true
	}
	return false
}

// Intervals returns the intervals in betwen notes, duplicated notes are removed.
func (c *Chord) Intervals() []uint {
	if c.isIntervalCacheValid() {
		return c.KeyIntervals
	}

	keys := c.Keys
	// remove duplicate notes (including those played on different octaves)
	seenKeys := map[int]bool{}
	keys = []int{}
	var pitch int
	for _, k := range c.Keys {
		pitch = k % 12
		if _, ok := seenKeys[pitch]; !ok {
			seenKeys[pitch] = true
			keys = append(keys, pitch)
		}
	}

	c.KeyIntervals = keyIntervals(keys)
	c.intervalKeyCache = make([]int, len(c.Keys))
	copy(c.intervalKeyCache, c.Keys)
	return c.KeyIntervals
}

// cache validation check
func (c *Chord) isIntervalCacheValid() bool {
	if c == nil || len(c.KeyIntervals) == 0 || len(c.intervalKeyCache) != len(c.Keys) {
		return false
	}

	for i, k := range c.intervalKeyCache {
		if k != c.Keys[i] {
			return false
		}
	}

	return true
}

func keyIntervals(keys []int) []uint {
	var lastKey int
	var pitch, lastPitch int
	intervals := []uint{}
	for i, k := range keys {
		if i == 0 {
			lastKey = k
			continue
		}
		pitch, lastPitch = k%12, lastKey%12
		interval := pitch - lastPitch
		if interval < 0 {
			interval = (k % 12) + 12 - lastKey%12
		}
		if interval > 8 {
			interval = -(pitch - lastPitch)
		}
		intervals = append(intervals, uint(interval))
		lastKey = k
	}
	return intervals
}
