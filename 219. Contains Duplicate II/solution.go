package main

import "math"

func main() {

}

func containsNearbyDuplicate(nums []int, k int) bool {

	n := len(nums)
	if n < 2 || k == 0 {
		return false
	}

	nmap := make(map[int]int)

	for i := 0; i < n; i++ {
		j, exists := nmap[nums[i]]
		if exists && math.Abs(float64(i-j)) <= float64(k) {
			return true
		}
		nmap[nums[i]] = i
	}
	return false
}
