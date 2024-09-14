package main

func main() {
	_ = strStr("HAYSTACKCYSTACKPP", "YSTAC")
}

func strStr(haystack string, needle string) int {

	//p, i, j := -1, 0, 0
	var i int16 = 0
	var j int16 = 0
	var p int = -1
	for int(j) < len(haystack) {
		if []rune(needle)[i] == []rune(haystack)[j] {
			if i == 0 {
				p = int(j)
			}
			i++
			j++
			if int(i) == len(needle) {
				return p
			}
		} else {
			if i > 0 {
				j = int16(p + 1)
			} else {
				j++
			}
			i = 0
			p = -1
		}
	}
	if int(i) != len(needle) {
		return -1
	}
	return p
}
