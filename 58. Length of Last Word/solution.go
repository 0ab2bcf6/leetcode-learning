package main

func main() {

}

func lengthOfLastWord(s string) int {
	rslice := []rune(s)
	rfound := false
	count := 0
	for i := len(s) - 1; i > -1; i-- {
		if (rslice[i] >= 65 && rslice[i] <= 90) || (rslice[i] >= 97 && rslice[i] <= 122) {
			rfound = true
			count++
		}
		if rfound && !((rslice[i] >= 65 && rslice[i] <= 90) || (rslice[i] >= 97 && rslice[i] <= 122)) {
			return count
		}
	}
	return count
}
