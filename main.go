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
	res, err := intToRoman()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
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

func newRomanConverter() RomanConverter {
	return roman.RomanConverterFast{}
}
