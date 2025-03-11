package main

import (
	"bufio"
	"fmt"
	"log"
	"njcejvbehrvbehr/roman"
	"os"
	"strconv"
)

type RomanConverter interface {
	IntToRoman(int) (string, error)
	RomanToInt(string) (int, error)
}

func main() {
	scan := bufio.NewScanner(os.Stdin)

	fmt.Println("Roman to Int (R) or int to roman (I)?")
	scan.Scan()
	scanned := scan.Text()
	switch scanned {
	case "I":
		res, err := intToRoman()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res)
	case "R":
		res, err := romanToInt()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res)
	default:
		fmt.Printf("Invalid operation. Enter either a 'R' or 'I'. Your input was: '%s'", scanned)
	}

}

func intToRoman() (string, error) {
	scan := bufio.NewScanner(os.Stdin)

	fmt.Println("enter a number")
	scan.Scan()
	scanned := scan.Text()

	inputAsInt, err := strconv.Atoi(scanned)
	if err != nil {
		return "", err
	}

	converter := newRomanConverter()
	return converter.IntToRoman(inputAsInt)
}

func romanToInt() (int, error) {
	scan := bufio.NewScanner(os.Stdin)

	fmt.Println("enter a Roman numeral")
	scan.Scan()
	scanned := scan.Text()

	converter := newRomanConverter()
	return converter.RomanToInt(scanned)
}

func newRomanConverter() RomanConverter {
	return roman.RomanConverterFast{}
}
