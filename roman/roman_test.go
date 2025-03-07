package roman_test

import (
	"njcejvbehrvbehr/roman"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	testcases := []testcase{
		{
			input:       -1,
			expectError: true,
		},
		{
			input:          1,
			expectedOutput: "I",
		},
	}

	for _, tc := range testcases {
		res, err := roman.IntToRoman(tc.input)
		if err != nil && !tc.expectError {
			t.Error(err)
			continue
		}
		if tc.expectError && err == nil {
			t.Errorf("expected error for input '%d', but got none", tc.input)
			continue
		}

		if res != tc.expectedOutput {
			t.Errorf("expected output '%s' for input '%d', but got '%s'", tc.expectedOutput, tc.input, res)
		}
	}
}

type testcase struct {
	input          int
	expectedOutput string
	expectError    bool
}
