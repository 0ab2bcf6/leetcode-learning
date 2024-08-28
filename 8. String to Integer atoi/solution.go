package main

import "math"

func main() {

}
func myAtoi(s string) int {

	result := 0
	negative := false
	pos := 0

	// loop until valid starting character is found
	for i, r := range s {
		// ignore leading whitespace
		if r == 32 {
			continue
		} else if r == 43 || r == 45 {
			if r == 45 {
				negative = true
			}
			pos = i + 1
			break
		} else if isDigit(r) {
			pos = i
			break
		} else if !isDigit(r) {
			return 0
		}
	}

	srunes := []rune(s)
	for i := pos; i < len(srunes); i++ {
		r := srunes[i]

		if !isDigit(r) {
			return getResult(result, negative)
		} else if result == 0 && r == 48 {
			continue
		} else if result != 0 && r == 48 {
			if result*10 > math.MaxInt32 {
				if negative {
					return math.MinInt32
				}
				return math.MaxInt32
			}
			result *= 10
		} else {
			if result*10+int(r-'0') > math.MaxInt32 {
				if negative {
					return math.MinInt32
				}
				return math.MaxInt32
			}
			result *= 10
			result += int(r - '0')
		}
	}
	return getResult(result, negative)
}

func getResult(result int, neg bool) int {

	if neg {
		return -result
	}
	return result
}

func isDigit(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	}
	return false
}
