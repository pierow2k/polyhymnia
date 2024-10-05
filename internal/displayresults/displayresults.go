// Package displayresults provides functions to display the results
// of a query.
package displayresults

import (
	"fmt"
	"strings"

	"github.com/pierow2k/polyhymnia/internal/datamuseapi"
)

// DisplayOptions holds flags for controlling which parts of the result
// are displayed.
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

// ToMetadataString generates the Metadata string used in the query based
// on the flags set in DisplayOptions. It appends corresponding
// abbreviations to the existing metadata string if their associated flags
// are set.
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

// displayIfNotZero prints an integer value with its label
// if the value is greater than zero.
func displayIfNotZero(label string, value int) {
	if value > 0 {
		fmt.Printf("\t%s: %d\n", label, value)
	}
}

// displayIfNotZeroFloat prints a float value with its label
// if the value is greater than zero.
func displayIfNotZeroFloat(label string, value float64) {
	if value > 0 {
		fmt.Printf("\t%s: %.6f\n", label, value)
	}
}

// displayIfNotEmpty prints a string value with its label
// if the value is non-empty.
func displayIfNotEmpty(label, value string) {
	if value != "" {
		fmt.Printf("\t%s: %s\n", label, value)
	}
}

// displayDefinitions prints the definitions of a resultif the flag is set.
func displayDefinitions(result datamuseapi.APIResponse, show bool) {
	if show && len(result.Definitions) > 0 {
		for _, def := range result.Definitions {
			fmt.Printf("\tDefinition: %s\n", def)
		}
	}
}

// displayPOS prints the parts of speech of a result if the flag is set.
func displayPOS(result datamuseapi.APIResponse, show bool) {
	if show && len(result.Tags) > 0 {
		for _, tag := range result.Tags {
			fmt.Printf("\tPart of Speech: %s\n", tag)
		}
	}
}

// displaySingleResult prints the details of a single APIResponse based on
// the provided DisplayOptions.
func displaySingleResult(result datamuseapi.APIResponse, options DisplayOptions) {
	fmt.Printf("%s\n", result.Word)

	if options.ShowScore {
		displayIfNotZero("Score", result.Score)
	}

	if options.ShowSyllables {
		displayIfNotZero("Num Syllables", result.NumSyllables)
	}

	if options.ShowPronunciation {
		displayIfNotEmpty("Pronunciation", result.Pronunciation)
	}

	if options.ShowFrequency {
		displayIfNotZeroFloat("Frequency", result.Frequency)
	}

	displayPOS(result, options.ShowPOS)
	displayDefinitions(result, options.ShowDefinitions)

	fmt.Println() // Line break between results
}

// DisplayResults processes and displays a list of APIResponse objects
// based on the flags set in DisplayOptions. It prints a summary of the
// query (such as the URL and result count) if requested, followed by the
// detailed results.
func DisplayResults(results []datamuseapi.APIResponse, options DisplayOptions) {
	// Display query URL if requested.
	if options.ShowQueryURL && len(results) > 0 {
		fmt.Printf("Datamuse API URL: %s\n", results[0].QueryURL)
	}

	// Display count of results if requested.
	if options.ShowCountFlag {
		fmt.Printf("Number of Results: %d\n", len(results))
	}

	// Display each result based on the options provided.
	for _, result := range results {
		displaySingleResult(result, options)
	}
}
