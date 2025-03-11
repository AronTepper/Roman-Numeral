package roman

import (
	"fmt"
	"regexp"
	"strings"
)

func (c RomanConverterFast) RomanToInt(numeral string) (int, error) {
	res := 0
	prevValue := 0
	repeatCount := 0

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
		val, ok := romanMap[v]

		if !ok {
			return 0, fmt.Errorf("cannot convert invalid Roman numeral, it contains: '%v'", string(v))
		}

		if val < prevValue {
			switch {
			case (prevValue == 5 || prevValue == 10) && val != 1:
				return 0, fmt.Errorf("invalid Roman numeral, it contains an illegal sequence: '%v'", string(v))
			case (prevValue == 50 || prevValue == 100) && val != 10:
				return 0, fmt.Errorf("invalid Roman numeral, it contains an illegal sequence: '%v'", string(v))
			case (prevValue == 500 || prevValue == 1000) && val != 100:
				return 0, fmt.Errorf("invalid Roman numeral, it contains an illegal sequence: '%v'", string(v))
			}
			res -= val
		} else {
			res += val
		}

		if val != prevValue {
			repeatCount = 0
		}

		if val == prevValue {
			if val != 1 && val != 10 && val != 100 && val != 1000 {
				return 0, fmt.Errorf("invalid Roman numeral, it contains 2 sequential: '%v'", string(v))
			}
			repeatCount++
		}
		if v != 'M' && repeatCount > 2 {
			return 0, fmt.Errorf("invalid Roman numeral, it contains more than 3 sequential: '%v'", string(v))
		}

		prevValue = val
	}

	return res, nil
}

type RomanConverterRegEx struct{}

func (c RomanConverterRegEx) RomanToIntRegex(s string) (int, error) {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	s = strings.ToUpper(s)

	if !isValidRoman(s) {
		return 0, fmt.Errorf("invalid input: '%s'", s)
	}

	res := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		c := rune(s[i])
		val := romanMap[c]
		if val < prevValue {
			res -= val
		} else {
			res += val
		}
		prevValue = val
	}
	return res, nil
}

func isValidRoman(s string) bool {
	validRomanPattern := `^M*(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`
	matched, _ := regexp.MatchString(validRomanPattern, s)
	return matched
}
