package main

import (
	"sort"
)

func main() {
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	_ = threeSum([]int{-1, 0, 1, 2, -1, -4})
}

func threeSum(nums []int) [][]int {

	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				duplicate := false
				sol := []int{nums[i], nums[l], nums[r]}
				for _, arr := range result {
					if isEqual(sol, arr) {
						duplicate = true
						break
					}
				}
				if !duplicate {
					result = append(result, sol)
				}
				l++
				r--
			} else if nums[l]+nums[r]+nums[i] > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return result
}

func isEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	// a and b are sorted
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
