package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{2, 1, 2, 0, 1}))
}

func maxProfit(prices []int) int {

	if len(prices) == 1 {
		return 0
	}

	if len(prices) == 2 {
		if prices[0] > prices[1] {
			return 0
		}
		return prices[1] - prices[0]
	}

	maxprofit := 0
	j := 0
	for i := 1; i < len(prices); i++ {
		if prices[j] < prices[i] {
			maxprofit = int(math.Max(float64(maxprofit), float64(prices[i]-prices[j])))
		} else {
			j = i
		}
	}

	return maxprofit
}
