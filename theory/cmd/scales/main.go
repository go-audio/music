package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-audio/music/theory"
)

/*
	Demo command line that returns the notes in the provided scale, for instance:
	$ scales b min
	Notes in B Natural Minor: [B C# D E F# G A]
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need to pass the tonic/root key to get the notes for.")
		os.Exit(1)
	}
	tonic := os.Args[1]
	scaleName := "Major"
	if len(os.Args) > 2 {
		scaleName = strings.Join(os.Args[2:], " ")
	}
	scaleName = strings.Title(scaleName)
	if scaleName == "Minor" || scaleName == "Min" {
		scaleName = "Natural Minor"
	}
	tonic = strings.ToUpper(tonic)
	scale, ok := theory.ScaleDefMap[theory.ScaleName(scaleName)]
	if !ok {
		fmt.Printf("Couldn't find the scale you asked for (%s), pick one of the following:\n", scaleName)
		for _, s := range theory.ScaleDefs {
			fmt.Printf("\t%s\n", s.Name)
		}
		os.Exit(1)
	}
	noteInts, noteNames := theory.ScaleNotes(tonic, scale.Name)
	fmt.Printf("Notes in %s %s: %v\n", tonic, scale.Name, noteNames)
	fmt.Printf("Key indexes in %s %s: %v\n", tonic, scale.Name, noteInts)
	fmt.Println()

	// var notes = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	// for _, tonic := range notes {
	// 	ints, _ := theory.ScaleNotes(tonic, theory.ScaleName(scaleName))
	// 	fmt.Printf("%s major: %v\n", tonic, ints)
	// }
}
