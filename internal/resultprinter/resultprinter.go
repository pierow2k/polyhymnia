// Package resultprinter provides functions to format and display query
// results from the Datamuse API based on customizable display options.
package resultprinter

import (
	"fmt"
	"strings"

	"github.com/pierow2k/polyhymnia/internal/datamuseapi"
)

// DisplayOptions contains flags to control which parts of the query
// result should be displayed, such as definitions, frequency, or
// part of speech.
type DisplayOptions struct {
	ShowCountFlag     bool
	ShowDefinitions   bool
	ShowFrequency     bool
	ShowPOS           bool
	ShowPronunciation bool
	ShowQueryURL      bool
	ShowScore         bool
	ShowSyllables     bool
}

// ToMetadataString generates a metadata string for the query by appending
// relevant abbreviations based on the flags set in DisplayOptions. If a
// flag is enabled, the corresponding abbreviation is added to the metadata.
func (opts *DisplayOptions) ToMetadataString(existingMd string) string {
	if strings.ToLower(existingMd) == "none" {
		existingMd = ""
	}

	var builder strings.Builder

	builder.WriteString(existingMd)

	addToMd := func(abbreviation string) {
		if !strings.Contains(existingMd, abbreviation) {
			builder.WriteString(abbreviation)
		}
	}

	if opts.ShowDefinitions {
		addToMd("d")
	}

	if opts.ShowFrequency {
		addToMd("f")
	}

	if opts.ShowPOS {
		addToMd("p")
	}

	if opts.ShowPronunciation {
		addToMd("r")
	}

	if opts.ShowSyllables {
		addToMd("s")
	}

	return builder.String()
}

// printIntIfNotZero prints a labeled integer value
// if it is greater than zero.
func printIntIfNotZero(label string, value int) {
	if value > 0 {
		fmt.Printf("\t%s: %d\n", label, value)
	}
}

// printFloatIfNotZero prints a labeled float value
// if it is greater than zero.
func printFloatIfNotZero(label string, value float64) {
	if value > 0 {
		fmt.Printf("\t%s: %.6f\n", label, value)
	}
}

// printStringIfNotEmpty prints a labeled string value
// if it is non-empty.
func printStringIfNotEmpty(label, value string) {
	if value != "" {
		fmt.Printf("\t%s: %s\n", label, value)
	}
}

// printDefinitions prints the definitions of the query result if
// the show flag is set and definitions are available.
func printDefinitions(result datamuseapi.APIResponse, show bool) {
	if show && len(result.Definitions) > 0 {
		for _, def := range result.Definitions {
			fmt.Printf("\tDefinition: %s\n", def)
		}
	}
}

// printPartOfSpeech prints the parts of speech (POS) for the query result
// if the ShowPOS flag is enabled.
func printPartOfSpeech(result datamuseapi.APIResponse, show bool) {
	if show && len(result.Tags) > 0 {
		for _, tag := range result.Tags {
			fmt.Printf("\tPart of Speech: %s\n", tag)
		}
	}
}

// printResultDetails prints the details of a single APIResponse,
// including score, syllables, pronunciation, etc., based on the
// flags set in DisplayOptions.
func printResultDetails(result datamuseapi.APIResponse, options DisplayOptions) {
	fmt.Printf("%s\n", result.Word)

	if options.ShowScore {
		printIntIfNotZero("Score", result.Score)
	}

	if options.ShowSyllables {
		printIntIfNotZero("Num Syllables", result.NumSyllables)
	}

	if options.ShowPronunciation {
		printStringIfNotEmpty("Pronunciation", result.Pronunciation)
	}

	if options.ShowFrequency {
		printFloatIfNotZero("Frequency", result.Frequency)
	}

	printPartOfSpeech(result, options.ShowPOS)
	printDefinitions(result, options.ShowDefinitions)

	fmt.Println() // Line break between results.
}

// PrintResults processes and displays a list of APIResponse objects
// according to the flags set in DisplayOptions. It prints the API query
// URL and result count if those options are enabled, followed by the
// detailed results for each response.
func PrintResults(results []datamuseapi.APIResponse, options DisplayOptions) {
	// Display the API query URL if the ShowQueryURL flag is enabled.
	if options.ShowQueryURL && len(results) > 0 {
		fmt.Printf("Datamuse API URL: %s\n", results[0].QueryURL)
	}

	// Display the count of results if the ShowCountFlag is enabled.
	if options.ShowCountFlag {
		fmt.Printf("Number of Results: %d\n", len(results))
	}

	// Display each result based on the options provided.
	for _, result := range results {
		printResultDetails(result, options)
	}
}
