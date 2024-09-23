package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "martgana"))
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	n := len(s)
	sslice, tslice := []rune(s), []rune(t)
	rmap := make(map[rune]int)
	for i := 0; i < n; i++ {
		count, exists := rmap[sslice[i]]
		if !exists {
			rmap[sslice[i]] = 0
		}
		count++
		rmap[sslice[i]] = count

		count, exists = rmap[tslice[i]]
		if !exists {
			rmap[tslice[i]] = 0
		}
		count--
		rmap[tslice[i]] = count

	}
	for _, v := range rmap {
		if v != 0 {
			return false
		}
	}
	return true
}
