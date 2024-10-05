// Package datamuseapi provides functions to query the Datamuse API and
// handle its responses.
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

// RequestTimeout defines the maximum duration allowed for API requests,
// set to ten seconds.
const RequestTimeout = 10 * time.Second

// QueryParams defines the parameters used for making a query to the
// Datamuse API.
type QueryParams struct {
	Lc         string   `url:"lc,omitempty"`     // Left context
	Max        int      `url:"max,omitempty"`    // Maximum number of results to return
	Md         string   `url:"md,omitempty"`     // Metadata flags
	Ml         bool     `url:"-"`                // If true, performs a "means like" search.
	Sl         bool     `url:"-"`                // If true, performs a  "Sounds like" search.
	Sp         bool     `url:"-"`                // If true, performs a "Spelled like" search.
	Qe         string   `url:"qe,omitempty"`     // Query echo
	Rc         string   `url:"rc,omitempty"`     // Right context
	RelCode    []string `url:"rel_,omitempty"`   // Related word constraints with a code.
	Topics     []string `url:"topics,omitempty"` // Topic words (space or comma delimited).
	V          string   `url:"v,omitempty"`      // Identifier for the vocabulary to use.
	SearchTerm string   `url:"search_term"`      // SearchTerm represents the word or phrase to be searched in the API.
}

// APIResponse holds the response data returned by the Datamuse API.
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

// buildQueryURL constructs the URL for querying the Datamuse API based
// on QueryParams.
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

	// When using rel, append SearchTerm to each related word constraint.
	for _, rel := range q.RelCode {
		appendParam("rel_"+rel, q.SearchTerm)
	}

	if q.Max > 0 {
		appendParam("max", strconv.Itoa(q.Max))
	}

	queryURL := strings.TrimSuffix(builder.String(), "&")

	return queryURL
}

// parseAPIResponse processes the raw API response and extracts additional
// fields like pronunciation, frequency, and the original query URL.
func parseAPIResponse(rawResponse []APIResponse, queryURL string) []APIResponse {
	for i, result := range rawResponse {
		result.Pronunciation, result.Tags = extractPronunciation(result.Tags)
		result.Frequency, result.Tags = extractFrequency(result.Tags)
		result.QueryURL = queryURL // Set the query URL in the response.
		rawResponse[i] = result
	}

	return rawResponse
}

// extractTagPrefix extracts a specific prefix from a list of tags and
// returns the extracted value along with the updated tags.
func extractTagPrefix(tags []string, prefix string) (string, []string) {
	var (
		extracted   string
		updatedTags []string
	)

	for _, tag := range tags {
		if strings.HasPrefix(tag, prefix) {
			extracted = strings.TrimPrefix(tag, prefix)
		} else {
			updatedTags = append(updatedTags, tag)
		}
	}

	return extracted, updatedTags
}

// extractPronunciation extracts the pronunciation from the given tags
// by using the "pron:" prefix.
func extractPronunciation(tags []string) (string, []string) {
	return extractTagPrefix(tags, "pron:")
}

// extractFrequency extracts the frequency value from the given tags
// by using the "f:" prefix and returns it as a float.
func extractFrequency(tags []string) (float64, []string) {
	frequencyStr, updatedTags := extractTagPrefix(tags, "f:")
	frequency, _ := strconv.ParseFloat(frequencyStr, 64)

	return frequency, updatedTags
}

// QueryAPI sends a request to the Datamuse API based on the provided
// query parameters and returns the parsed API response or an error.
func QueryAPI(queryParams QueryParams, client *http.Client) ([]APIResponse, error) {
	// Build the query URL.
	queryURL := queryParams.buildQueryURL()

	// Parse and validate the queryURL.
	parsedURL, err := url.Parse(queryURL)
	if err != nil || !parsedURL.IsAbs() {
		return nil, fmt.Errorf("%w: invalid URL: %s", ErrAPIError, queryURL)
	}

	// Create a request context with a timeout to limit the duration of
	// the API request.
	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	defer cancel()

	// Build an HTTP GET request with the specified context and query URL.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, queryURL, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to create request: %w", ErrAPIError, err)
	}

	// Send the request using the provided client.
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to send request: %w", ErrAPIError, err)
	}

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Printf("error closing response body: %v\n", closeErr)
		}
	}()

	// Check if the response status is '200 OK' to ensure a successful
	// request.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: unexpected response code: %d", ErrAPIError, resp.StatusCode)
	}

	// Read the response body into a byte slice for further processing.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to read response body: %w", ErrAPIError, err)
	}

	// Unmarshal the response body from JSON into the APIResponse slice.
	var apiResponses []APIResponse

	err = json.Unmarshal(body, &apiResponses)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to parse JSON: %w", ErrAPIError, err)
	}

	// Parse the API response to extract pronunciation, frequency, and
	// the query URL for each result.
	parsedResponses := parseAPIResponse(apiResponses, queryURL)

	return parsedResponses, nil
}
