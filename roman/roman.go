package roman

import "fmt"

type pivot struct {
	value   int
	numeral string
}

func IntToRoman(number int) (string, error) {
	var piv = []pivot{
		{2000, "MM"},
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
