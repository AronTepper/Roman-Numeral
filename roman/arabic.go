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
	topValue := 0
	// usedValues := new([7]rune)

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

		val, ok := romanMap[v]

		if !ok {
			return 0, fmt.Errorf("cannot convert invalid Roman numeral, it contains: '%v'", v)
		}

		if v != 'M' && val == prevValue && repeatCount == 4 {
			return 0, fmt.Errorf("invalid Roman numeral, it contains more than 3 sequential: '%v'", v)
		}

		if val > topValue {
			res += val
			topValue = val
			continue
		}
		if val > prevValue {
			res -= val
		}
		if val == prevValue {
			repeatCount = 0
			res += val
		}
		if val == prevValue {
			repeatCount++
		}

		prevValue = val
	}

	fmt.Println("hoi")
	return res, nil
}

func (c RomanConverterFast) RomanToIntRegex(s string) (int, error) {
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
	validRomanPattern := `^(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`
	matched, _ := regexp.MatchString(validRomanPattern, s)
	return matched
}
