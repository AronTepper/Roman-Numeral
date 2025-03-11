package roman_test

import (
	"math/rand"
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

func BenchmarkRomanToInt(b *testing.B) {
	converter := roman.RomanConverterFast{}

	romanNumerals := make([]string, 3999)
	for i := 1; i <= 3999; i++ {
		v, _ := converter.IntToRoman(i)
		romanNumerals[i-1] = v
	}

	for i := 0; i < b.N; i++ {
		randomIndex := rand.Intn(len(romanNumerals))
		_, err := converter.RomanToInt(romanNumerals[randomIndex])
		if err != nil {
			b.Fatal(err)
		}
	}
}


// func BenchmarkRomanToIntRegEx(b *testing.B) {
// 	converter := roman.RomanConverterRegEx{}

// 	romanNumerals := make([]string, 3999)
// 	for i := 1; i <= 3999; i++ {
// 		v, _ := converter.RomanToIntRegex(i)
// 		romanNumerals[i-1] = v
// 	}

// 	for i := 0; i < b.N; i++ {
// 		randomIndex := rand.Intn(len(romanNumerals))
// 		_, err := converter.RomanToInt(romanNumerals[randomIndex])
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }



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
