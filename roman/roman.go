package roman

import (
	"fmt"
	"strings"
)

type RomanConverterFast struct{}

func (c RomanConverterFast) IntToRoman(number int) (string, error) {
	if number < 0 {
		return "", fmt.Errorf("invalid number")
	}

	var rv string
	if number >= 1000 {
		rv += strings.Repeat("M", number/1000)
	}
	if number >= 100 {
		rv += checkFunkyRomans((number%1000)/100, "C", "D", "M")
	}
	if number >= 10 {
		rv += checkFunkyRomans((number%100)/10, "X", "L", "C")
	}
	if number > 0 {
		rv += checkFunkyRomans(number%10, "I", "V", "X")
	}

	return rv, nil // fmt.Errorf("not implemented")
}

func checkFunkyRomans(number int, one, five, ten string) string {
	if number < 4 {
		return strings.Repeat(one, number)
	}
	if number == 4 {
		return one + five
	}
	if number >= 5 && number < 9 {
		return five + strings.Repeat(one, number-5)
	}
	return one + ten
}

type RomanConverterOld struct{}

func (c RomanConverterOld) IntToRoman(number int) (string, error) {
	type pivot struct {
		value   int
		numeral string
	}

	var piv = []pivot{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var res = ""

	if number < 0 {
		return "", fmt.Errorf("invalid number")
	}

	for _, val := range piv {
		for number >= val.value {
			res += val.numeral
			number -= val.value
		}
	}

	return res, nil
}
