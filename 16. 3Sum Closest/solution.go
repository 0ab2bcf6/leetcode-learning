package main

import (
	"math"
	"sort"
)

func main() {
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	_ = threeSumClosest([]int{-1, 0, 1, 2, -1, -4}, 1)
}

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	globalmin := math.MaxInt

	for i := 1; i < len(nums); i++ {

	}
	return globalmin
}

func _oldthreeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	globalmin := math.MaxInt

	a, b := 0, len(nums)-1
	for i := 1; i < len(nums); i++ {

		currmin := math.MaxInt

		for math.Abs(float64(nums[a]+nums[i]+nums[b]-target)) < float64(currmin) {

			currsol := nums[a] + nums[i] + nums[b] - target
			if currsol == 0 {
				return 0
			}

			if currsol < currmin {
				currmin = currsol
			}

			if currsol > 0 && math.Abs(float64(nums[a]+nums[i]+nums[b-1]-target)) < math.Abs(float64(currsol)) && (b-1) != i {
				b--
			} else if currsol < 0 && math.Abs(float64(nums[a+1]+nums[i]+nums[b]-target)) < math.Abs(float64(currsol)) && (a+1) != i {
				a++
			} else {
				break
			}
		}

		if currmin < globalmin {
			globalmin = currmin
		}
	}
	return globalmin
}
