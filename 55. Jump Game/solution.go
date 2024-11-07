package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 0, 0}))
}

func canJump(nums []int) bool {

	if len(nums) == 1 {
		return true
	}
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if max <= 0 {
			return false
		}
		max--
		if nums[i] >= max {
			max = nums[i]
		}
	}
	return true
}
