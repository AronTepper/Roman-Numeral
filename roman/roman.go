package roman

import (
	"fmt"
	"strings"
)

type pivot struct {
	value   int
	numeral string
}

func IntToRoman(number int) (string, error) {
	if number < 0 {
		return "", fmt.Errorf("invalid number")
	}

	thousands := number / 1000
	hundreds := (number % 1000) / 100
	tens := (number % 100) / 10
	ones := number % 10

	res := ""
	res += strings.Repeat("M", thousands)
	res += checkFunkyRomans(hundreds, "C", "D", "M")
	res += checkFunkyRomans(tens, "X", "L", "C")
	res += checkFunkyRomans(ones, "I", "V", "X")

	return res, nil // fmt.Errorf("not implemented")
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

func IntToRomanOld(number int) (string, error) {
	var piv = []pivot{
		// {2000, "MM"},
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
		fmt.Println(val)
		for number >= val.value {
			res += val.numeral
			number -= val.value
		}
	}

	return res, nil // fmt.Errorf("not implemented")
}
