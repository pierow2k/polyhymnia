// Package datamuseapi queries the Datamuse API.
package datamuseapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const requestTimeout = 10 * time.Second // Timeout in ten seconds

// QueryParams struct defines the parameters used to query the API.
type QueryParams struct {
	Lc         string   `url:"lc,omitempty"`     // Left context
	Max        int      `url:"max,omitempty"`    // Maximum number of results to return
	Md         string   `url:"md,omitempty"`     // Metadata flags
	Ml         bool     `url:"-"`                // Indicates if "Means like" search should be performed
	Sl         bool     `url:"-"`                // Indicates if "Sounds like" search should be performed
	Sp         bool     `url:"-"`                // Indicates if "Spelled like" search should be performed
	Qe         string   `url:"qe,omitempty"`     // Query echo
	Rc         string   `url:"rc,omitempty"`     // Right context
	RelCode    []string `url:"rel_,omitempty"`   // Related word constraints with a code
	Topics     []string `url:"topics,omitempty"` // Topic words, space or comma delimited
	V          string   `url:"v,omitempty"`      // Identifier for the vocabulary to use
	SearchTerm string   `url:"search_term"`      // Search term for the query
}

// APIResponse struct is used to hold the API response data.
//
//nolint:tagliatelle
type APIResponse struct {
	Word          string   `json:"word"`                // Vocabulary entry
	Score         int      `json:"score"`               // Word rank
	NumSyllables  int      `json:"numSyllables"`        // Syllable count
	Tags          []string `json:"tags"`                // Parts of speech (e.g., noun, verb, adj)
	Definitions   []string `json:"defs"`                // Definitions
	Pronunciation string   `json:"pron,omitempty"`      // Pronunciation (extracted from tags)
	Frequency     float64  `json:"frequency,omitempty"` // Word frequency (extracted from tags)
	QueryURL      string   `json:"queryURL,omitempty"`  // The API query URL
}

// ErrAPIError is a package-level error for API failures.
var ErrAPIError = errors.New("datamuse api error")

// buildQueryURL builds the query URL.
func (q *QueryParams) buildQueryURL() string {
	baseURL := "https://api.datamuse.com/words"

	var builder strings.Builder

	builder.WriteString(baseURL + "?")

	appendParam := func(key, value string) {
		if value != "" {
			builder.WriteString(fmt.Sprintf("%s=%s&", url.QueryEscape(key), url.QueryEscape(value)))
		}
	}

	// Append the correct search parameter based on which boolean is true.
	if q.Ml {
		appendParam("ml", q.SearchTerm)
	}

	if q.Sl {
		appendParam("sl", q.SearchTerm)
	}

	if q.Sp {
		appendParam("sp", q.SearchTerm)
	}

	appendParam("v", q.V)
	appendParam("lc", q.Lc)
	appendParam("rc", q.Rc)
	appendParam("md", q.Md)
	appendParam("qe", q.Qe)

	if len(q.Topics) > 0 {
		appendParam("topics", strings.Join(q.Topics, ","))
	}

	// When using rel, append SearchTerm to each related word constraint
	for _, rel := range q.RelCode {
		appendParam("rel_"+rel, q.SearchTerm)
	}

	if q.Max > 0 {
		appendParam("max", strconv.Itoa(q.Max))
	}

	queryURL := strings.TrimSuffix(builder.String(), "&")

	return queryURL
}

// parseAPIResponse processes the raw API response and calls helper
// functions to extract the pronunciation and frequency.
func parseAPIResponse(rawResponse []APIResponse, queryURL string) []APIResponse {
	for i, result := range rawResponse {
		result.Pronunciation, result.Tags = extractPronunciation(result.Tags)
		result.Frequency, result.Tags = extractFrequency(result.Tags)
		result.QueryURL = queryURL // Set the query URL in the response
		rawResponse[i] = result
	}

	return rawResponse
}

func extractPronunciation(tags []string) (string, []string) {
	var pronunciation string

	var updatedTags []string

	for _, tag := range tags {
		if strings.HasPrefix(tag, "pron:") {
			pronunciation = strings.TrimPrefix(tag, "pron:")
		} else {
			updatedTags = append(updatedTags, tag)
		}
	}

	return pronunciation, updatedTags
}

func extractFrequency(tags []string) (float64, []string) {
	var frequency float64

	var updatedTags []string

	for _, tag := range tags {
		if strings.HasPrefix(tag, "f:") {
			freqStr := strings.TrimPrefix(tag, "f:")
			if freq, err := strconv.ParseFloat(freqStr, 64); err == nil {
				frequency = freq
			}
		} else {
			updatedTags = append(updatedTags, tag)
		}
	}

	return frequency, updatedTags
}

// QueryAPI sends a request to the Datamuse API with the given query
// parameters and returns the response.
func QueryAPI(queryParams QueryParams) ([]APIResponse, error) {
	// Build the query URL
	queryURL := queryParams.buildQueryURL()

	// Parse and validate the queryURL
	parsedURL, err := url.Parse(queryURL)
	if err != nil || !parsedURL.IsAbs() {
		return nil, fmt.Errorf("%w: invalid URL: %s", ErrAPIError, queryURL)
	}

	// Define an HTTP client with a timeout
	client := &http.Client{Timeout: requestTimeout}

	// Set a timeout for the request context
	ctx := context.Background()

	// Create an HTTP GET request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, queryURL, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to create request: %w", ErrAPIError, err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to send request: %w", ErrAPIError, err)
	}

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Printf("error closing response body: %v\n", closeErr)
		}
	}()

	// Check if the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: unexpected response code: %d", ErrAPIError, resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to read response body: %w", ErrAPIError, err)
	}

	// Parse the JSON response
	var apiResponses []APIResponse

	err = json.Unmarshal(body, &apiResponses)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to parse JSON: %w", ErrAPIError, err)
	}

	// Parse the response to extract Pronunciation, Frequency, and include query URL
	parsedResponses := parseAPIResponse(apiResponses, queryURL)

	return parsedResponses, nil
}
