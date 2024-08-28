package main

import "sort"

func main() {
}

func NaivefirstMissingPositive(nums []int) int {

	for i := 1; i < len(nums)+1; i++ {
		smallest := false
		for _, num := range nums {
			smallest = smallest || (i == num)
		}
		if !smallest {
			return i
		}
	}
	return len(nums) + 1
}

func ImprovedfirstMissingPositive(nums []int) int {

	sort.Ints(nums)

	smallest := 1
	for _, num := range nums {
		if num == smallest {
			smallest++
		}
	}
	return smallest
}
