package main

import "math"

func main() {
}

func jump(nums []int) int {

	i := 0

	for j := nums[i]; j > 0; j-- {
		minSteps(nums[i+j])
	}
}

func minSteps(idx int, nums []int) int {
	if nums[idx] == 0 {
		return math.MaxInt32
	}
	for j := nums[idx]; j > 0; j-- {

	}
}
