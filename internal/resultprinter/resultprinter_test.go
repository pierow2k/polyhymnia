// Package resultprinter_test provides tests for the resultprinter package.
package resultprinter_test

import (
	"testing"

	"github.com/pierow2k/polyhymnia/internal/resultprinter"
)

// TestToMetadataString tests the ToMetadataString function in various
// scenarios.
//
//nolint:funlen
func TestToMetadataString(t *testing.T) {
	t.Parallel()

	// Table-based test cases
	tests := []struct {
		name       string
		opts       resultprinter.DisplayOptions
		existingMd string
		expectedMd string
	}{
		{
			name:       "No flags set, empty existingMd",
			opts:       resultprinter.DisplayOptions{},
			existingMd: "",
			expectedMd: "",
		},
		{
			name:       "No flags set, existingMd set to 'none'",
			opts:       resultprinter.DisplayOptions{},
			existingMd: "none",
			expectedMd: "",
		},
		{
			name: "ShowDefinitions flag set, empty existingMd",
			opts: resultprinter.DisplayOptions{
				ShowDefinitions: true,
			},
			existingMd: "",
			expectedMd: "d",
		},
		{
			name: "ShowFrequency flag set, existingMd already contains 'd'",
			opts: resultprinter.DisplayOptions{
				ShowFrequency: true,
			},
			existingMd: "d",
			expectedMd: "df",
		},
		{
			name: "Multiple flags set, empty existingMd",
			opts: resultprinter.DisplayOptions{
				ShowDefinitions:   true,
				ShowFrequency:     true,
				ShowPOS:           true,
				ShowPronunciation: true,
				ShowSyllables:     true,
			},
			existingMd: "",
			expectedMd: "dfprs",
		},
		{
			name: "Multiple flags set, existingMd contains some abbreviations",
			opts: resultprinter.DisplayOptions{
				ShowDefinitions: true,
				ShowFrequency:   true,
				ShowPOS:         true,
			},
			existingMd: "f",
			expectedMd: "fdp",
		},
		{
			name:       "No new flags, existingMd remains the same",
			opts:       resultprinter.DisplayOptions{},
			existingMd: "prs",
			expectedMd: "prs",
		},
	}

	// Iterate over each test case
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			// Call the function with the test inputs
			result := testCase.opts.ToMetadataString(testCase.existingMd)

			// Compare the result with the expected output
			if result != testCase.expectedMd {
				t.Errorf("ToMetadataString() = %v, want %v", result, testCase.expectedMd)
			}
		})
	}
}
