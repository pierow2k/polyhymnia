package datamuseapi_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/pierow2k/polyhymnia/internal/datamuseapi"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest
func TestQueryAPI_NoResults(t *testing.T) {
	// Start HTTP mocking
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock URL and empty response
	httpmock.RegisterResponder("GET", "https://api.datamuse.com/words?sp=thststrtrnsnthng",
		httpmock.NewStringResponder(200, `[]`))

	// Define query parameters
	queryParams := datamuseapi.QueryParams{
		Sp:         true,
		SearchTerm: "thststrtrnsnthng",
	}

	// Execute QueryAPI
	client := &http.Client{}
	results, err := datamuseapi.QueryAPI(queryParams, client)

	// Assertions
	require.NoError(t, err)
	require.Empty(t, results)
}

//nolint:paralleltest
func TestQueryAPI_SuccessfulSpelledLike(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock URL and response for spelled-like search
	//nolint:lll
	httpmock.RegisterResponder("GET", "https://api.datamuse.com/words?sp=pyro%2A&max=3",
		httpmock.NewStringResponder(200, `[{"word":"pyro","score":522},{"word":"pyrometry","score":339},{"word":"pyrolysis","score":276}]`))

	// Define query parameters
	queryParams := datamuseapi.QueryParams{
		Sp:         true,
		Max:        3,
		SearchTerm: "pyro*",
	}

	// Execute QueryAPI
	client := &http.Client{}
	results, err := datamuseapi.QueryAPI(queryParams, client)

	// Assertions
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, "pyro", results[0].Word)
	require.Equal(t, "pyrometry", results[1].Word)
	require.Equal(t, "pyrolysis", results[2].Word)
}

//nolint:paralleltest
func TestQueryAPI_SuccessfulSoundsLike(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock URL and response for sounds-like search
	//nolint:lll
	httpmock.RegisterResponder("GET", "https://api.datamuse.com/words?sl=pie+row&max=3",
		httpmock.NewStringResponder(200, `[{"word":"pyro","score":100,"numSyllables":2},{"word":"pyro-","score":100,"numSyllables":2},{"word":"pairo","score":95,"numSyllables":2}]`))

	// Define query parameters
	queryParams := datamuseapi.QueryParams{
		Sl:         true,
		Max:        3,
		SearchTerm: "pie row",
	}

	// Execute QueryAPI
	client := &http.Client{}
	results, err := datamuseapi.QueryAPI(queryParams, client)

	// Assertions
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, "pyro", results[0].Word)
	require.Equal(t, 2, results[0].NumSyllables)
}

//nolint:paralleltest
func TestQueryAPI_SuccessfulMeansLike(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock URL and response for means-like search
	//nolint:lll
	httpmock.RegisterResponder("GET", "https://api.datamuse.com/words?ml=pyro&max=3",
		httpmock.NewStringResponder(200, `[{"word":"metol","score":19931501,"tags":["n","results_type:primary_rel"]},{"word":"pyrotechnics","score":10028771,"tags":["n"]},{"word":"pyrotechnic","score":10028770,"tags":["adj"]}]`))

	// Define query parameters
	queryParams := datamuseapi.QueryParams{
		Ml:         true,
		Max:        3,
		SearchTerm: "pyro",
	}

	// Execute QueryAPI
	client := &http.Client{}
	results, err := datamuseapi.QueryAPI(queryParams, client)

	// Assertions
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, "metol", results[0].Word)
	require.Equal(t, "pyrotechnics", results[1].Word)
}

//nolint:paralleltest
func TestQueryAPI_SuccessfulRelatedAntonym(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock URL and response for antonym search
	//nolint:lll
	httpmock.RegisterResponder("GET", "https://api.datamuse.com/words?rel_ant=fire&max=3",
		httpmock.NewStringResponder(200, `[{"word":"employ","score":5358},{"word":"engage","score":3196},{"word":"hire","score":1200}]`))

	// Define query parameters
	queryParams := datamuseapi.QueryParams{
		RelCode:    []string{"ant"},
		Max:        3,
		SearchTerm: "fire",
	}

	// Execute QueryAPI
	client := &http.Client{}
	results, err := datamuseapi.QueryAPI(queryParams, client)

	// Assertions
	require.NoError(t, err)
	require.Len(t, results, 3)
	require.Equal(t, "employ", results[0].Word)
	require.Equal(t, "engage", results[1].Word)
	require.Equal(t, "hire", results[2].Word)
}

//nolint:paralleltest,funlen
func TestQueryAPI_SuccessfulMetadataSearch(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock URL and response for metadata search
	//nolint:lll
	httpmock.RegisterResponder("GET", "https://api.datamuse.com/words?ml=burn&md=dfprs&max=2",
		httpmock.NewStringResponder(200, `[
			{
				"word": "burn off",
				"score": 30050032,
				"numSyllables": 2,
				"tags": ["syn","v","results_type:primary_rel","pron:B ER1 N AO1 F ","f:0.204818"],
				"defs": [
					"v\t(intransitive) To dissipate as the result of heat.",
					"v\t(transitive) To cause to dissipate by applying heat.",
					"v\t(transitive, intransitive, oil) To dispose of (unusable explosive natural gas from an oil well) by burning it as it emerges from the well.",
					"v\t(intransitive, rail transport, of an axle bearing) To fail due to overheating.",
					"v\t(transitive, intransitive, television) To fill (low-value air time) with programming not suitable for its original purpose.",
					"v\t(transitive) To expend energy resulting from metabolizing food.",
					"v\t(transitive) To use up a resource in a nonproductive manner.",
					"v\t(rugby) Cause to waste energy."
				]
			},
			{
				"word": "burn down",
				"score": 30049985,
				"numSyllables": 2,
				"tags": ["syn","v","pron:B ER1 N D AW1 N ","f:0.416113"],
				"defs": [
					"v\t(transitive) To cause (a structure) to burn to nothing.",
					"v\t(intransitive, of a structure) To burn completely, so that nothing remains."
				]
			}
		]`))

	// Define query parameters
	queryParams := datamuseapi.QueryParams{
		Ml:         true,
		Md:         "dfprs", // Metadata flags
		Max:        2,
		SearchTerm: "burn",
	}

	// Execute QueryAPI
	client := &http.Client{}
	results, err := datamuseapi.QueryAPI(queryParams, client)

	// Assertions
	require.NoError(t, err)
	require.Len(t, results, 2)

	// Check first result
	require.Equal(t, "burn off", results[0].Word)
	require.Equal(t, 2, results[0].NumSyllables)
	require.Contains(t, results[0].Tags, "syn")
	require.Contains(t, results[0].Tags, "v")
	require.Equal(t, "B ER1 N AO1 F ", results[0].Pronunciation)
	//nolint:testifylint
	require.Equal(t, 0.204818, results[0].Frequency)

	// Check second result
	require.Equal(t, "burn down", results[1].Word)
	require.Equal(t, 2, results[1].NumSyllables)
	require.Equal(t, "B ER1 N D AW1 N ", results[1].Pronunciation)
	//nolint:testifylint
	require.Equal(t, 0.416113, results[1].Frequency)
}

// func TestQueryAPI_InvalidURL(t *testing.T) {
// 	// Define invalid query parameters
// 	queryParams := datamuseapi.QueryParams{
// 		SearchTerm: " ", // This will generate an invalid URL
// 		Sp:         true,
// 	}

// 	// Execute QueryAPI
// 	client := &http.Client{}
// 	results, err := datamuseapi.QueryAPI(queryParams, client)

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Contains(t, err.Error(), "invalid URL")
// 	assert.Nil(t, results)
// }

//nolint:paralleltest
func TestQueryAPI_RequestSendFailure(t *testing.T) {
	// Mock http.NewRequestWithContext to return an error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterNoResponder(func(_ *http.Request) (*http.Response, error) {
		//nolint:err113
		return nil, errors.New("request creation failed")
	})

	// Define query parameters
	queryParams := datamuseapi.QueryParams{
		Sp:         true,
		SearchTerm: "validterm",
	}

	// Execute QueryAPI
	client := &http.Client{}
	results, err := datamuseapi.QueryAPI(queryParams, client)

	// Assertions
	require.Error(t, err)
	require.ErrorContains(t, err, "failed to send request")
	require.Nil(t, results)
}

// TestQueryAPI_HTTPClientFailure mocks a network error using httpmock
// to simulate a client failure (e.g., network timeout or connection
// refusal).
//
//nolint:paralleltest
func TestQueryAPI_HTTPClientFailure(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Simulate a network error
	//nolint:err113
	httpmock.RegisterResponder("GET", "https://api.datamuse.com/words?sp=pyro",
		httpmock.NewErrorResponder(errors.New("network error")))

	// Define query parameters
	queryParams := datamuseapi.QueryParams{
		Sp:         true,
		SearchTerm: "pyro",
	}

	// Execute QueryAPI
	client := &http.Client{}
	results, err := datamuseapi.QueryAPI(queryParams, client)

	// Assertions
	require.Error(t, err)
	require.ErrorContains(t, err, "network error")
	require.Nil(t, results)
}

// TestQueryAPI_Non200Response mocks a non-200 response (e.g., 404 Not
// Found) to test how the function handles unexpected status codes.
//
//nolint:paralleltest
func TestQueryAPI_Non200Response(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Simulate a 404 error response
	httpmock.RegisterResponder("GET", "https://api.datamuse.com/words?sp=pyro",
		httpmock.NewStringResponder(404, "Not Found"))

	// Define query parameters
	queryParams := datamuseapi.QueryParams{
		Sp:         true,
		SearchTerm: "pyro",
	}

	// Execute QueryAPI
	client := &http.Client{}
	results, err := datamuseapi.QueryAPI(queryParams, client)

	// Assertions
	require.Error(t, err)
	require.ErrorContains(t, err, "unexpected response code: 404")
	require.Nil(t, results)
}

//nolint:paralleltest
func TestQueryAPI_ResponseBodyReadError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Simulate a broken body (forces an error when reading the body)
	httpmock.RegisterResponder("GET", "https://api.datamuse.com/words?sp=pyro",
		httpmock.NewStringResponder(200, `invalid body`))

	// Define query parameters
	queryParams := datamuseapi.QueryParams{
		Sp:         true,
		SearchTerm: "pyro",
	}

	// Execute QueryAPI
	client := &http.Client{}
	results, err := datamuseapi.QueryAPI(queryParams, client)

	// Assertions
	require.Error(t, err)
	require.ErrorContains(t, err, "failed to parse JSON")
	require.Nil(t, results)
}

//nolint:paralleltest
func TestQueryAPI_InvalidJSONResponse(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Simulate invalid JSON response
	httpmock.RegisterResponder("GET", "https://api.datamuse.com/words?sp=pyro",
		httpmock.NewStringResponder(200, `{"invalid json"}`))

	// Define query parameters
	queryParams := datamuseapi.QueryParams{
		Sp:         true,
		SearchTerm: "pyro",
	}

	// Execute QueryAPI
	client := &http.Client{}
	results, err := datamuseapi.QueryAPI(queryParams, client)

	// Assertions
	require.Error(t, err)
	require.ErrorContains(t, err, "failed to parse JSON")
	require.Nil(t, results)
}
