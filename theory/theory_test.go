package theory

import (
	"github.com/go-audio/midi"
)

func keyNames(keys []int) []string {
	names := make([]string, len(keys))
	for i, k := range keys {
		names[i] = midi.NoteToName(k)
	}
	return names
}
