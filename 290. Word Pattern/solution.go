package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat dog giraffe"))

}

func wordPattern(pattern string, s string) bool {
	sarr := split(s, ' ')
	parr := []rune(pattern)

	if len(sarr) != len(parr) {
		return false
	}

	rmap := make(map[rune]string)
	smap := make(map[string]rune)

	for i := 0; i < len(parr); i++ {
		word, rexists := rmap[parr[i]]
		char, wexists := smap[sarr[i]]

		if rexists && wexists && (word != sarr[i] || char != parr[i]) {
			return false
		} else if (rexists && !wexists) || (!rexists && wexists) {
			return false
		}

		rmap[parr[i]] = sarr[i]
		smap[sarr[i]] = parr[i]

	}
	return true
}

func split(s string, seperator rune) []string {
	f := func(r rune) bool {
		return r == seperator
	}
	return strings.FieldsFunc(s, f)
}
