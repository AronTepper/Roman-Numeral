package roman_test

import (
	"njcejvbehrvbehr/roman"
	"testing"
)

func TestIntToRomanFast(t *testing.T) {
	runIntToRomanTestcases(t, roman.RomanConverterFast{})
}

func TestIntToRomanOld(t *testing.T) {
	runIntToRomanTestcases(t, roman.RomanConverterOld{})
}

func BenchmarkIntToRomanFast(b *testing.B) {
	converter := roman.RomanConverterFast{}

	for i := 0; b.Loop(); i++ {
		_, err := converter.IntToRoman(i)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIntToRomanOld(b *testing.B) {
	converter := roman.RomanConverterOld{}

	for i := 0; b.Loop(); i++ {
		_, err := converter.IntToRoman(i)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Define an unexported interface here as well so we can test different
// implementations with the same func.
type converter interface {
	IntToRoman(int) (string, error)
}

func runIntToRomanTestcases(t *testing.T, converter converter) {
	for _, tc := range testcases {
		res, err := converter.IntToRoman(tc.input)
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

var testcases = []testcase{
	{
		input:       -1,
		expectError: true,
	},
	{
		input:          0,
		expectedOutput: "",
	},
	{
		input:          1,
		expectedOutput: "I",
	},
	{
		input:          4,
		expectedOutput: "IV",
	},
	{
		input:          9,
		expectedOutput: "IX",
	},
	{
		input:          40,
		expectedOutput: "XL",
	},
	{
		input:          90,
		expectedOutput: "XC",
	},
	{
		input:          400,
		expectedOutput: "CD",
	},
	{
		input:          900,
		expectedOutput: "CM",
	},
	{
		input:          58,
		expectedOutput: "LVIII",
	},
	{
		input:          1994,
		expectedOutput: "MCMXCIV",
	},
	{
		input:          3999,
		expectedOutput: "MMMCMXCIX",
	},
	// Maybe?
	// {
	// 	input:       4000,
	// 	expectError: true,
	// },
}
