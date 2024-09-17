package main

import "fmt"

func main() {
	_ = removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}

func removeDuplicates(nums []int) int {
	k := 0
	for _, num := range nums {
		if k == 0 || nums[k-1] != num {
			nums[k] = num
			k++
		}
	}
	return k
}

func removeDuplicatesNotInPlace(nums []int) int {
	k := 0
	curr := -101
	n := len(nums)
	for _, num := range nums {
		if num != curr {
			curr = num
			k++
			nums = append(nums, curr)
		}
	}
	nums = nums[n:]
	fmt.Println(nums)
	return k
}
