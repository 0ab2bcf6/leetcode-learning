package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{1, 3, 4, 5, 6, 7, 9, 10, 12, 13}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

func summaryRanges(nums []int) []string {
	n := len(nums)
	var result []string
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if j == i {
				if j == len(nums)-1 || nums[j] != nums[j+1]-1 {
					result = append(result, fmt.Sprintf("%d", nums[j]))
					i = j
					break
				}
			} else {
				if j == len(nums)-1 || nums[j] != nums[j+1]-1 {
					result = append(result, fmt.Sprintf("%d->%d", nums[i], nums[j]))
					i = j
					break
				}
			}
		}
	}
	return result
}
