package main

func main() {
}

// TODO improve
func stupidsearch(nums []int, target int) int {

	if nums[0] == target || nums[len(nums)-1] == target {
		if nums[0] == target {
			return 0
		}
		return len(nums) - 1
	}

	if nums[0] > target && nums[len(nums)-1] > target {
		for i := 0; i < len(nums); i++ {
			if nums[i] == target {
				return i
			}
		}
	} else if nums[0] < target && nums[len(nums)-1] < target {
		for i := len(nums) - 1; i > 0; i-- {
			if nums[i] == target {
				return i
			}
		}
	} else {
		// normal sorted list
		l, r := 0, len(nums)-1
		for l <= r {
			m := (l + r) / 2
			if nums[m] < target {
				l = m + 1
			} else if nums[m] > target {
				r = m - 1
			} else {
				return m
			}
		}
	}
	return -1
}
