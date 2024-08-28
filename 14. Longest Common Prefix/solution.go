package main

import (
	"fmt"
	"sort"
)

func main() {
	longestCommonPrefix([]string{"abab", "aba", "abc"})
}

func longestCommonPrefix(strs []string) string {

	//TODO
	fmt.Println(strs)
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	for _, c := range strs {
		fmt.Println(c)
	}

	return ""
}
