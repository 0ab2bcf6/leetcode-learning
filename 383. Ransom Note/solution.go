package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine) {
		return false
	}

	rmap := make(map[rune]int)
	for _, r := range magazine {
		c, exists := rmap[r]
		if !exists {
			rmap[r] = 1
		} else {
			c++
			rmap[r] = c
		}
	}
	for _, r := range ransomNote {
		c, exists := rmap[r]
		if exists {
			c--
			if c < 0 {
				return false
			}
			rmap[r] = c
		} else {
			return false
		}
	}
	return true
}
