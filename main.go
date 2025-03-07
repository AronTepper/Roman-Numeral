package main

import (
	"fmt"
	"njcejvbehrvbehr/arabic"
	"njcejvbehrvbehr/roman"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(roman.IntToRoman(1))
	fmt.Println(arabic.RomanToInt("I"))
}
