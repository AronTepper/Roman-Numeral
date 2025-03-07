package main

import (
	"fmt"
	"strconv"
	"strings"
)

func romanToInt(s string) (int, error) {
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
	total := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		c := rune(s[i])
		value, ok := romanMap[c]
		if !ok {
			return 0, fmt.Errorf("ongeldig karakter '%c' in Romeins cijfer", c)
		}
		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}
	return total, nil
}

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"MM", "CM", "DD", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			result.WriteString(romans[i])
		}
	}
	return result.String()
}

func main() {
	var input string
	fmt.Println("Enter value:")

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Wrong input:", err)
		return
	}

	// Probeer de invoer als Arabisch getal te parsen
	if num, err := strconv.Atoi(input); err == nil {
		// Controleer of het getal binnen het gebruikelijke bereik (1-3999) ligt
		// if num <= 0 || num > 3999 {
		// 	fmt.Println("Geef een getal tussen 1 en 3999 op.")
		// 	return
		// }
		roman := intToRoman(num)
		fmt.Printf("Arabisch: %d -> Romeins: %s\n", num, roman)
	} else {
		// Anders gaan we ervan uit dat de invoer een Romeins cijfer is.
		num, err := romanToInt(input)
		if err != nil {
			fmt.Println("Ongeldig Romeins cijfer:", err)
			return
		}
		fmt.Printf("Romeins: %s -> Arabisch: %d\n", strings.ToUpper(input), num)
	}
}
