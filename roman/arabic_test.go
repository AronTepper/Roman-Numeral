package roman_test

import (
	"njcejvbehrvbehr/roman"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	converter := roman.RomanConverterFast{}
	for _, tc := range testcasesArab {
		res, err := converter.RomanToInt(tc.input)
		if err != nil && !tc.expectError {
			t.Error(err)
			continue
		}
		if tc.expectError && err == nil {
			t.Errorf("expected error for input '%s', but got none", tc.input)
			continue
		}

		if res != tc.expectedOutput {
			t.Errorf("expected output '%d' for input '%s', but got '%d'", tc.expectedOutput, tc.input, res)
		}
	}
}



var testcasesArab = []testcaseArab{
	{
		input:          "",
		expectedOutput: 0,
	},
	{
		input:          "I",
		expectedOutput: 1,
	},
	{
		input:          "IV",
		expectedOutput: 4,
	},
	{
		input:          "IX",
		expectedOutput: 9,
	},
	{
		input:          "X",
		expectedOutput: 10,
	},
	{
		input:          "XL",
		expectedOutput: 40,
	},
	{
		input:          "L",
		expectedOutput: 50,
	},
	{
		input:          "XC",
		expectedOutput: 90,
	},
	{
		input:          "C",
		expectedOutput: 100,
	},
	{
		input:          "CD",
		expectedOutput: 400,
	},
	{
		input:          "D",
		expectedOutput: 500,
	},
	{
		input:          "CM",
		expectedOutput: 900,
	},
	{
		input:          "M",
		expectedOutput: 1000,
	},
	{
		input:          "MMXXIII",
		expectedOutput: 2023,
	},
	{
		input:       "invalid",
		expectError: true,
	},
	{
		input:       "IIII",
		expectError: true,
	},
	{
		input:       "VV",
		expectError: true,
	},
	{
		input:       "XXXX",
		expectError: true,
	},
	{
		input:       "LL",
		expectError: true,
	},
	{
		input:       "CCCC",
		expectError: true,
	},
	{
		input:       "DD",
		expectError: true,
	},
	{
		input:       "IC",
		expectError: true,
	},
	{
		input:       "IL",
		expectError: true,
	},
	{
		input:       "VX",
		expectError: true,
	},
	{
		input:       "XD",
		expectError: true,
	},
	{
		input:       "LC",
		expectError: true,
	},
	{
		input:       "DM",
		expectError: true,
	},
	// Maybe?
	// {
	// 	input:       "MMMM",
	// 	expectError: true,
	// },

}

type testcaseArab struct {
	input          string
	expectedOutput int
	expectError    bool
}
