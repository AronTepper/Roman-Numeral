package main

import (
	"bufio"
	"fmt"
	"log"
	"njcejvbehrvbehr/roman"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("enter a number")
	scan.Scan()
	scanned := scan.Text()

	// fmt.Println(scanned)
	conv, err := strconv.Atoi(scanned)
	if err != nil {
		log.Fatal(err)
		return
	}
	res, err := roman.IntToRoman(conv)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res)
	// fmt.Println(arabic.RomanToInt("I"))
}
