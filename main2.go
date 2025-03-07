package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func CountWords(s string) map[string]int {
	resp := map[string]int{}
	words := strings.Fields(s)

	for _, V := range words {
		resp[V]++
	}
	// for idx := range words {
	// 	resp[words[idx]]++
	// }
	// for _, V := range strings.Split(s, " ") {
	// 	resp[V]++
	// }
	return resp

}

func main() {
	wc.Test(CountWords)
}
