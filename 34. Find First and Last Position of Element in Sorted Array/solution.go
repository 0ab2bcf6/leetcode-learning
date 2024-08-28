package main

func main() {
}

func searchRange(nums []int, target int) []int {
	// idea: binary search
	l, r := 0, len(nums)-1
	i, j := len(nums), 0
	found := false
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			i, j = m, m
			found = true
			break
		}
	}
	if found {
		for (i - 1) >= 0 {
			if nums[i-1] == target {
				i--
			} else {
				break
			}
		}
		for (j + 1) < len(nums) {
			if nums[j+1] == target {
				j++
			} else {
				break
			}
		}
		return []int{i, j}
	}
	return []int{-1, -1}
}
