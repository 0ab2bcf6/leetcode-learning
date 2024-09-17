package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	rmap := make(map[rune]rune)
	vmap := make(map[rune]rune)

	sarr, tarr := []rune(s), []rune(t)
	for i := 0; i < len(sarr); i++ {
		r, diff1 := mapRune(rmap, sarr[i], tarr[i])
		_, diff2 := mapRune(vmap, tarr[i], sarr[i])
		if r != tarr[i] || diff1 || diff2 {
			return false
		}
	}
	return true
}

func mapRune(rmap map[rune]rune, key rune, val rune) (rune, bool) {
	r, exists := rmap[key]
	// fmt.Println(key, val, exists)
	if !exists {
		rmap[key] = val
		return val, false
	} else if exists && r != val {
		// fmt.Println("reached")
		return r, true
	}
	return r, false
}
