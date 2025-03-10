package roman

import (
	"fmt"
	"strings"
	"unicode"
)

func (c RomanConverterFast) RomanToInt(numeral string) (int, error) {
	res := 0
	lastValue := 0

	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	s := strings.ToUpper(numeral)

	for i := range s {
		v := rune(s[len(s)-1-i])

		fmt.Println(s, i, v)

		if romanMap[v] > lastValue {
			res += romanMap[v]
			lastValue = romanMap[v]
		}

		// switch {
		// case v == 'M' && lastValue > 1000:
		// 	res += 1000
		// 	lastValue = 1000
		// case v == 'M' && lastValue > 1000:
		// 	res += 1000
		// 	lastValue = 1000
		// case v == 'M' && lastValue > 1000:
		// 	res += 1000
		// 	lastValue = 1000

		// case v == 'M' && lastValue > 1000:
		// 	res += 1000
		// 	lastValue = 1000

		// case v == 'M' && lastValue > 1000:
		// 	res += 1000
		// 	lastValue = 1000
		// }
	}

	fmt.Println("hoi")
	return res, nil
}

func isValidRomanNumeral(s string) error {
	if areDigits(s) {
		return fmt.Errorf("cannot convert invalid Roman numeral, it contains digits: '%s'", s)
	}
	if !areLetters(s) {
		return fmt.Errorf("cannot convert invalid Roman numeral, it does not contain letters only: '%s'", s)
	}

	return nil
}

func areDigits(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func areLetters(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}
