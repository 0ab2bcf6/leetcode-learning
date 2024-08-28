package main

func main() {

}

// TODO optimize
func lengthOfLongestSubstring(s string) int {
	var currMax int = 0
	var baseLength int = len(s)

	for i := 0; i < baseLength; i++ {
		charMap := make(map[byte]int)
		charMap[s[i]] = i

		for j := i + 1; j < baseLength; j++ {
			_, ok := charMap[s[j]]
			if ok {
				if len(charMap) > currMax {
					currMax = len(charMap)
				}
				break
			} else {
				charMap[s[j]] = j
			}
		}
		if len(charMap) > currMax {
			currMax = len(charMap)
		}
	}
	return currMax
}
