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
}

func main() {
	scan := bufio.NewScanner(os.Stdin)

	fmt.Println("enter a number")
	scan.Scan()
	scanned := scan.Text()

	conv, err := strconv.Atoi(scanned)

	if err != nil {
		log.Fatal(err)
		return
	}
	converter := newConverter()
	res, err := converter.IntToRoman(conv)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res)
}

func newConverter() RomanConverter {
	return roman.RomanConverterOld{}
}
