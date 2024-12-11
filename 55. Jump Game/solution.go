package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {

	if len(nums) == 1 {
		return true
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		max--
		if i == len(nums)-1 {
			return true
		}
		if nums[i] >= max {
			max = nums[i]
		}
		if max < 0 {
			return false
		}
	}
	return true
}
