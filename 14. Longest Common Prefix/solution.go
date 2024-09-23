package main

import "sort"

func main() {
	longestCommonPrefix([]string{"abab", "aba", "abc"})
}

func longestCommonPrefix(strs []string) string {

	var result []rune
	// var current []rune
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	shortest, longest := strs[0], strs[len(strs)-1]
	sarr, larr := []rune(shortest), []rune(longest)
	for i := 0; i < min(len(shortest), len(longest)); i++ {
		if sarr[i] != larr[i] {
			return string(result)
		}
		result = append(result, sarr[i])
	}

	return string(result)
}
