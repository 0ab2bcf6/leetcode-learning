package main

import "math"

func main() {}

func maxProfit(prices []int) int {
	min, max := math.MaxInt, math.MinInt
	j := 0
	for i := len(prices) - 1; i > 0; i-- {
		if i > j {
			break
		}
		if min > prices[j] && i != j {
			min = prices[j]
		}
		if max < prices[i] && i != j {
			max = prices[i]
		}
		j++
	}

	if max != math.MinInt && min != math.MaxInt && max-min > 0 {
		return max - min
	}
	return 0
}
