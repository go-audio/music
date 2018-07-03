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
	// KeyIntervals are the half steps between each key, in most caes, you want to use Intervals().
	KeyIntervals     []uint
	intervalKeyCache []int

	sortedRelativeKeys []int
	// rootKey is the chord key of the chord. In the case of an inversion, the root key isn't the bass key.
	rootKey int
	// bassKey is the lowest key of the chord. It usually is the root, unless the chord is an inversion.
	bassKey int
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
	name = strings.ToLower(name)
	var chordRef *ChordDefinition
	for _, chordDef := range ChordDefs {
		if name == strings.ToLower(chordDef.Abbrev) {
			chordRef = chordDef
			break
		}
	}
	if chordRef == nil {
		// We probably have an inversion
		// FIXME: reculate the interval using the second note as the first
		return nil
	}
	chordRef.Root = string(root)
	rootInt := midi.KeyInt(chordRef.Root, 0)
	chord := &Chord{Keys: []int{rootInt}}
	var last int
	for _, interv := range chordRef.HalfSteps {
		last += int(interv)
		chord.Keys = append(chord.Keys, last)
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
	if c == nil {
		return ""
	}
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

// Root returns the root key of the chord
func (c *Chord) Root() int {
	if c == nil {
		return 0
	}
	if len(c.sortedRelativeKeys) == 0 && len(c.Keys) > 0 {
		c.sortedRelativeKeys = make([]int, len(c.Keys))
		sort.Ints(c.Keys)
		c.bassKey = c.Keys[0]
		copy(c.sortedRelativeKeys, c.Keys)
		for i := 0; i < len(c.sortedRelativeKeys); i++ {
			c.sortedRelativeKeys[i] = c.sortedRelativeKeys[i] % 12
		}
		sort.Ints(c.sortedRelativeKeys)
		c.rootKey = c.sortedRelativeKeys[0]
	}

	return c.rootKey
}

// Bass returns the bass key of the chord
func (c *Chord) Bass() int {
	if c == nil {
		return 0
	}

	if len(c.sortedRelativeKeys) == 0 && len(c.Keys) > 0 {
		c.Root()
	}

	return c.bassKey
}

// Def returns the matching chord definition with the root set if found.
// A chord definition with a name set to Unknown will be returned if no matches found.
func (c *Chord) Def() *ChordDefinition {
	if c == nil {
		return nil
	}
	// TODO: consider caching this result
	sort.Ints(c.Keys)
	retries := len(c.Keys)
	for retries > 0 {
		for _, chordDef := range ChordDefs {
			if c.Matches(chordDef) {
				return chordDef.WithRoot(midi.Notes[c.Root()])
			}
		}
		// we didn't find the chord, let's try to change the interval orders
		if len(c.Keys) < 2 {
			break
		}
		// fmt.Println("failed to find chord for", keysToNotes(c.Keys), intervals)
		c.Keys = append(c.Keys[1:], c.Keys[0])
		// fmt.Println("swapping keys", keysToNotes(c.Keys), intervals)
		retries--
	}

	// TODO: check for inversions (the root note isn't the bass note)
	// Sort the notes to be in order
	// check for a match on the new intervals
	// If there's a match, calculate the inversion (distance order from bass to the root note)
	// add a ^ to the chord per inversion #.

	return &ChordDefinition{Name: "Unknown"}
}

// Matches compares the current chord with the passed chord.
func (c *Chord) Matches(chordDef *ChordDefinition) bool {
	if reflect.DeepEqual(chordDef.HalfSteps, c.Intervals()) {
		return true
	}
	return false
}

// Intervals returns the intervals in betwen notes, duplicated notes are removed.
func (c *Chord) Intervals() []uint {
	if c == nil {
		return nil
	}
	if c.isIntervalCacheValid() {
		return c.KeyIntervals
	}

	c.Root()
	// remove duplicate notes (including those played on different octaves)
	seenKeys := map[int]bool{}
	keys := []int{}
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
