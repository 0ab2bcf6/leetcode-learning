package main

import (
	"sort"
)

func main() {
	// _ = maxTaxiEarnings(20, [][]int{{2, 5, 4}, {1, 5, 1}})
	// _ = maxTaxiEarnings(20, [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}})
	// _ = maxTaxiEarnings(20, [][]int{{2, 3, 4}, {4, 5, 1}})
}

// time limit exceeded
func maxTaxiEarnings(n int, rides [][]int) int64 {

	N := len(rides)
	result := make([][]int64, N)
	for i := range result {
		result[i] = make([]int64, N)
	}

	sort.Slice(rides, func(i, j int) bool {
		return rides[i][0] < rides[j][0]
	})

	var maxprofit int64 = 0
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			if i == 0 {
				result[i][j] = calcEarning(rides[j])
			} else {
				var maxi int64 = 0
				for k := 0; k < i; k++ {
					maxi = max(result[k][i-1], maxi)
				}

				if isReachable(rides[i-1], rides[j]) {
					result[i][j] = calcEarning(rides[j]) + maxi
				} else {
					result[i][j] = -1
				}
			}
			maxprofit = max(result[i][j], maxprofit)
		}
	}
	return maxprofit
}

func calcEarning(ride []int) int64 {
	return int64(ride[1] - ride[0] + ride[2])
}

func isReachable(a []int, b []int) bool {
	return b[0] >= a[1]
}
