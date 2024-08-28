package main

func longestPalindrome(s string) string {
	var result string = ""
	resLength := 0

	srunes := []rune(s)
	for pivot := range len(srunes) {

		// odd
		lpointer := pivot
		rpointer := pivot

		for lpointer >= 0 && rpointer < len(srunes) && srunes[lpointer] == srunes[rpointer] {
			if (rpointer - lpointer + 1) > resLength {
				result = string(srunes[lpointer : rpointer+1])
				resLength = rpointer - lpointer + 1
			}
			lpointer -= 1
			rpointer += 1
		}

		// even
		lpointer = pivot
		rpointer = pivot + 1
		for lpointer >= 0 && rpointer < len(srunes) && srunes[lpointer] == srunes[rpointer] {
			if (rpointer - lpointer + 1) > resLength {
				result = string(srunes[lpointer : rpointer+1])
				resLength = rpointer - lpointer + 1
			}
			lpointer -= 1
			rpointer += 1
		}

	}

	return result
}
