// Package cmd handles the command-line interface for Polyhymnia.
package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pierow2k/polyhymnia/internal/datamuseapi"
	"github.com/pierow2k/polyhymnia/internal/resultprinter"
	"github.com/spf13/cobra"
)

var (
	// Query parameters for the Datamuse API.
	queryParams datamuseapi.QueryParams
	// DisplayOptions struct to group all the display flags.
	displayOptions resultprinter.DisplayOptions
)

// init adds query and display option flags to RootCmd.
func init() {
	addQueryParamsFlags(RootCmd)
	addDisplayOptionsFlags(RootCmd)
}

// addQueryParamsFlags defines the query-related flags for API
// requests.
func addQueryParamsFlags(cmd *cobra.Command) {
	// Required flags
	cmd.Flags().BoolVarP(&queryParams.Ml, "means-like", "l", false, "Words with meaning similar to this string")
	cmd.Flags().BoolVarP(&queryParams.Sl, "sounds-like", "n", false, "Words that sound like this string")
	cmd.Flags().BoolVarP(&queryParams.Sp, "spelled-like", "t", false, "Words spelled like this string")
	cmd.Flags().StringArrayVar(&queryParams.RelCode, "related-word", []string{}, "Related word constraints")
	cmd.MarkFlagsOneRequired("means-like", "related-word", "sounds-like", "spelled-like")
	cmd.MarkFlagsMutuallyExclusive("means-like", "related-word", "sounds-like", "spelled-like")
	// Optional flags
	cmd.Flags().StringVar(&queryParams.V, "vocabulary", "", "Vocabulary identifier")
	cmd.Flags().StringArrayVar(&queryParams.Topics, "topics", []string{}, "Topics (comma-separated)")
	cmd.Flags().StringVar(&queryParams.Lc, "left-context", "", "Left context")
	cmd.Flags().StringVar(&queryParams.Rc, "right-context", "", "Right context")
	cmd.Flags().IntVar(&queryParams.Max, "max", 100, "Maximum number of results to return (1-1000)")
	cmd.Flags().StringVar(&queryParams.Md, "metadata", "", "Metadata flags (dfprs)")
}

// addDisplayOptionsFlags defines the flags that control output
// formatting.
func addDisplayOptionsFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&displayOptions.ShowCountFlag, "count", "c", false, "Show number of words returned by query")
	cmd.Flags().BoolVarP(&displayOptions.ShowDefinitions, "def", "d", false, "Include definitions in results")
	cmd.Flags().BoolVarP(&displayOptions.ShowFrequency, "freq", "f", false, "Include frequency in results")
	cmd.Flags().BoolVarP(&displayOptions.ShowPOS, "pos", "p", false, "Include parts of speech in results")
	cmd.Flags().BoolVarP(&displayOptions.ShowPronunciation, "pro", "r", false, "Include pronunciation in results")
	cmd.Flags().BoolVarP(&displayOptions.ShowScore, "score", "s", false, "Include score in results")
	cmd.Flags().BoolVarP(&displayOptions.ShowQueryURL, "show-query", "q", false, "Show the URL used for the query")
	cmd.Flags().BoolVarP(&displayOptions.ShowSyllables, "syl", "y", false, "Include syllables in results")
}

// setDisplayOptionsFromMetadata parses the metadata string and
// sets the appropriate display options.
func setDisplayOptionsFromMetadata(metadata string, opts *resultprinter.DisplayOptions) {
	flags := map[string]*bool{
		"d": &opts.ShowDefinitions,
		"f": &opts.ShowFrequency,
		"p": &opts.ShowPOS,
		"r": &opts.ShowPronunciation,
		"s": &opts.ShowSyllables,
	}

	metadata = strings.ToLower(metadata)
	for flag, option := range flags {
		if strings.Contains(metadata, flag) {
			*option = true
		}
	}
}

// runDatamuseQuery processes command-line flags, performs the API query,
// and displays the results.
func runDatamuseQuery(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no search term provided")
	}

	// Assign the first argument as SearchTerm.
	queryParams.SearchTerm = args[0]

	// Set displayOptions based on metadata flag (if provided).
	setDisplayOptionsFromMetadata(queryParams.Md, &displayOptions)

	// Construct the Metadata `Md` string from the display options.
	queryParams.Md = displayOptions.ToMetadataString(queryParams.Md)

	// Query the Datamuse API.
	client := &http.Client{
		Timeout: datamuseapi.RequestTimeout,
	}

	results, err := datamuseapi.QueryAPI(queryParams, client)
	if err != nil {
		return fmt.Errorf("error querying Datamuse API: %v", err)
	}

	if len(results) < 1 {
		fmt.Println("The search returned no results.")
	} else {
		// Display results using the resultprinter package.
		resultprinter.PrintResults(results, displayOptions)
	}

	return nil
}
