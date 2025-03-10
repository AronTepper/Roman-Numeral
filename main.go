package main

import (
	"bufio"
	"fmt"
	"log"
	"njcejvbehrvbehr/roman"
	"os"
	"strconv"
	"unicode"
)

type RomanConverter interface {
	IntToRoman(int) (string, error)
	RomanToInt(string) (int, error)
}


func main() {
	scan := bufio.NewScanner(os.Stdin)

	fmt.Println("enter a number")
	scan.Scan()
	scanned := scan.Text() 

	conv, err := strconv.Atoi(scanned)
	converter := newRomanConverter()

	if err != nil {

		log.Fatal(err)
		return
	}
	
	res, err := converter.IntToRoman(conv)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res)
}

func newRomanConverter() RomanConverter {
	return roman.RomanConverterFast{}
}

