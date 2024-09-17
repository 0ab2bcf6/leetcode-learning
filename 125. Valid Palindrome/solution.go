package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// fmt.Println(isPalindrome("race a car"))
	s := "A man, a plan, a canal: Panama"
	s = strings.ToLower(s)
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	// rslice := []rune(nonAlphanumericRegex.ReplaceAllString(s, ""))
	fmt.Println(nonAlphanumericRegex.ReplaceAllString(s, ""))

}

func isPalindrome(s string) bool {

	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	rslice := []rune(strings.ToLower(nonAlphanumericRegex.ReplaceAllString(s, "")))
	i, j := 0, len(rslice)-1

	for i < j {
		if rslice[i] != rslice[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isNumber(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	}
	return false
}
