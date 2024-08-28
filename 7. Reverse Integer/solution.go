package main

import "math"

func main() {

}

func reverse(x int) int {

	result := 0
	for x != 0 {
		nextDigit := x % 10
		if result < math.MinInt32 || result > math.MaxInt32 {
			return 0
		}
		result = result*10 + nextDigit
		x = x / 10
	}
	return result
}
