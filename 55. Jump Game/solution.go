package main

import "fmt"

func main() {
<<<<<<< HEAD
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
=======
	fmt.Println(canJump([]int{2, 0, 0}))
>>>>>>> 6e05abdcef42e69e0bc2ea7e5fb73adc21876928
}

func canJump(nums []int) bool {

	if len(nums) == 1 {
		return true
	}
<<<<<<< HEAD
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
=======
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if max <= 0 {
			return false
		}
		max--
		if nums[i] >= max {
			max = nums[i]
		}
>>>>>>> 6e05abdcef42e69e0bc2ea7e5fb73adc21876928
	}
	return true
}
